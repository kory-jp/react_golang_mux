package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kory-jp/react_golang_mux/api/usecase/transaction"

	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Todos   domain.Todos `json:"todos"`
	Todo    *domain.Todo `json:"todo"`
	SumPage float64      `json:"sumPage"`
}

func (res *Response) SetResp(status int, mess string, todos domain.Todos, todo *domain.Todo, sumPage float64) (resStr string) {
	response := &Response{status, mess, todos, todo, sumPage}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
			TodoTagRelationsRepository: &database.TodoTagRelationsRepository{
				SqlHandler: sqlHandler,
			},
			Transaction: transaction.SqlHandler(sqlHandler),
		},
	}
}

func GetUserId(r *http.Request) (userId int, err error) {
	session, err := controllers.Store.Get(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	if session.Values["userId"] == nil || session.Values["userId"] == 0 {
		return 0, err
	}

	userId = session.Values["userId"].(int)
	return userId, nil
}

// --- Todo新規追加 ----
func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "画像の容量が大きく保存できません", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "画像の取り込み失敗しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
	} else {
		defer file.Close()
		// 画像を保存するimgディレクトリが存在しない場合は作成する
		// devModeの画像保存先とtestModeの画像の保存先を"api/assets/dev/img"を共通化
		// devModeからtestModeでカレントディレクトリが異なるので、os.Getwdでそれぞれパスを指定する
		p, _ := os.Getwd()
		if string(p) == "/app/api" {
			fmt.Println("dev")
			err = os.MkdirAll("./assets/dev/img", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
				fmt.Fprintln(w, resStr)
				return
			}
		} else {
			err = os.MkdirAll("../../../assets/dev/img", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
				fmt.Fprintln(w, resStr)
				return
			}
		}

		// サーバー側に保存するために空ファイルを作成
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		var imageFilePath string
		if string(p) == "/app/api" {
			imageFilePath = "./assets/dev/img/" + uploadFileName
		} else {
			imageFilePath = "../../../assets/dev/img/" + uploadFileName
		}
		formatPath := filepath.Clean(imageFilePath)
		saveImage, err = os.Create(formatPath)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
		defer saveImage.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	todoType := new(domain.Todo)
	todoType.UserID = userId
	todoType.Title = r.Form.Get("title")
	todoType.Content = r.Form.Get("content")
	todoType.ImagePath = uploadFileName

	var tagIds []int
	ids := r.Form["tagIds"]
	if len(ids) != 0 {
		for _, v := range ids {
			toInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			tagIds = append(tagIds, toInt)
		}
	}

	mess, err := controller.Interactor.Add(*todoType, tagIds)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- Todo一覧取得 ---
func (controller *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	// URLから取得したいページ番目の情報
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	todos, sumPage, err := controller.Interactor.Todos(userId, page)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, "Todo一覧取得", todos, nil, sumPage)
	fmt.Fprintln(w, resStr)
}

// --- Todo詳細情報取得 ---
func (controller *TodoController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	todo, err := controller.Interactor.TodoByIdAndUserId(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "投稿詳細取得", nil, todo, 0)
	fmt.Fprintln(w, resStr)
}

func (controller *TodoController) TagSearch(w http.ResponseWriter, r *http.Request) {
	tagId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || tagId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	todos, sumPage, err := controller.Interactor.SearchTag(tagId, userId, page)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, "タグ検索成功", todos, nil, sumPage)
	fmt.Fprintln(w, resStr)
}

// ----- Todo更新 -----
func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "画像の容量が大きく保存できません", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "画像の取り込みに失敗しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
	} else {
		defer file.Close()

		p, _ := os.Getwd()
		if string(p) == "/app/api" {
			fmt.Println("dev")
			err = os.MkdirAll("./assets/dev/img", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
				fmt.Fprintln(w, resStr)
				return
			}
		} else {
			err = os.MkdirAll("../../../assets/dev/img", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
				fmt.Fprintln(w, resStr)
				return
			}
		}

		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		var imageFilePath string
		if string(p) == "/app/api" {
			imageFilePath = "./assets/dev/img/" + uploadFileName
		} else {
			imageFilePath = "../../../assets/dev/img/" + uploadFileName
		}
		formatPath := filepath.Clean(imageFilePath)
		saveImage, err = os.Create(formatPath)
		if err != nil {
			//  "サーバ側でファイル確保できませんでした
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
		defer saveImage.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			// アップロードしたファイルの書き込みに失敗しました。"
			fmt.Println(err)
			log.Println(err)
			resStr := new(Response).SetResp(400, "サーバーで障害が発生しました", nil, nil, 0)
			fmt.Fprintln(w, resStr)
			return
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	todoType := new(domain.Todo)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	todoType.ID = id
	todoType.UserID = userId
	todoType.Title = r.Form.Get("title")
	todoType.Content = r.Form.Get("content")
	// -------

	if r.Form.Get("imagePath") != "" {
		if uploadFileName != "" {
			// 画像変更
			todoType.ImagePath = uploadFileName
		} else {
			// 画像変更無し
			todoType.ImagePath = r.Form.Get("imagePath")
		}
	} else {
		if uploadFileName != "" {
			// 画像追加
			todoType.ImagePath = uploadFileName
		} else {
			// 画像無し、削除
			todoType.ImagePath = ""
		}
	}

	var tagIds []int
	ids := r.Form["tagIds"]
	fmt.Println("ids:", ids)
	if len(ids) != 0 {
		for _, v := range ids {
			toInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			tagIds = append(tagIds, toInt)
		}
	}

	// -------
	mess, err := controller.Interactor.UpdateTodo(*todoType, tagIds)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, mess.Message, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- Todo完了未完了を変更 ---
func (controller *TodoController) IsFinished(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	todoType := new(domain.Todo)
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	mess, err := controller.Interactor.IsFinishedTodo(id, *todoType, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- Todo削除 ---
func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	mess, err := controller.Interactor.DeleteTodo(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, mess.Message, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- Todo削除 + 一覧取得 ---
func (controller *TodoController) DeleteInIndex(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	todos, sumPage, mess, err := controller.Interactor.DeleteTodoInIndex(id, userId, page)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, todos, nil, sumPage)
	fmt.Fprintln(w, resStr)
}
