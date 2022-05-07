package controllers

import (
	"encoding/json"
	"errors"
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
	database "github.com/kory-jp/react_golang_mux/api/interfaces/database"
	todoTagRelations "github.com/kory-jp/react_golang_mux/api/interfaces/database/todo_tag_relations"
	todos "github.com/kory-jp/react_golang_mux/api/interfaces/database/todos"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todos"
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
			TodoRepository: &todos.TodoRepository{
				SqlHandler: sqlHandler,
			},
			TodoTagRelationsRepository: &todoTagRelations.TodoTagRelationsRepository{
				SqlHandler: sqlHandler,
			},
			Transaction: transaction.SqlHandler(sqlHandler),
		},
	}
}

// --- Todo新規追加 ----
func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	uploadFileName, err := MakeImagePath(r)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
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

	todoType := new(domain.Todo)
	todoType.UserID = userId
	todoType.Title = r.Form.Get("title")
	todoType.Content = r.Form.Get("content")
	todoType.ImagePath = uploadFileName
	todoType.Importance, err = strconv.Atoi(r.Form.Get("importance"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました!", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	todoType.Urgency, err = strconv.Atoi(r.Form.Get("urgency"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

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
// ---
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

// --- タグ,重要度,緊急度による検索 ---
// ---
func (controller *TodoController) Search(w http.ResponseWriter, r *http.Request) {
	tagId, err := strconv.Atoi(r.FormValue("tagId"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	importanceScore, err := strconv.Atoi(r.FormValue("importance"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	urgencyScore, err := strconv.Atoi(r.FormValue("urgency"))
	if err != nil {
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

	todos, sumPage, err := controller.Interactor.Search(tagId, importanceScore, urgencyScore, userId, page)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, "タグ検索成功", todos, nil, sumPage)
	fmt.Fprintln(w, resStr)
}

// ----- Todo更新 -----
// -----
func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	uploadFileName, err := MakeImagePath(r)
	if err != nil {
		resStr := new(Response).SetResp(400, err.Error(), nil, nil, 0)
		fmt.Fprintln(w, resStr)
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

	todoType.Importance, err = strconv.Atoi(r.Form.Get("importance"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	todoType.Urgency, err = strconv.Atoi(r.Form.Get("urgency"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

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
// ---
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

func MakeImagePath(r *http.Request) (uploadFileName string, err error) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		// resStr := new(Response).SetResp(400, "画像の容量が大きく保存できません", nil, nil, 0)
		err = errors.New("画像の容量が大きく保存できません")
		return "", err
	}
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
			return uploadFileName, nil
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("画像の取り込み失敗しました")
			return "", err
		}
	} else {
		defer file.Close()
		// 画像を保存するimgディレクトリが存在しない場合は作成する
		// devModeの画像保存先とtestModeの画像の保存先を"api/assets/dev/img"を共通化
		// devModeからtestModeでカレントディレクトリが異なるので、os.Getwdでそれぞれパスを指定する
		p, _ := os.Getwd()
		if string(p) == "/app/api" {
			// --- dev ----
			err = os.MkdirAll("./assets/dev/img", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				err = errors.New("サーバーで障害が発生しました")
				return "", err
			}
		} else {
			// --- test ---
			err = os.MkdirAll("../../../assets/dev/testImg", os.ModePerm)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				err = errors.New("サーバーで障害が発生しました")
				return "", err
			}
		}

		// サーバー側に保存するために空ファイルを作成
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		var imageFilePath string
		if string(p) == "/app/api" {
			imageFilePath = "./assets/dev/img/" + uploadFileName
		} else {
			imageFilePath = "../../../assets/dev/testImg/" + uploadFileName
		}
		formatPath := filepath.Clean(imageFilePath)
		saveImage, err = os.Create(formatPath)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("サーバーで障害が発生しました")
			return "", err

		}
		defer saveImage.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("サーバーで障害が発生しました")
			return "", err
		}
		fmt.Println("書き込んだByte数=>", size)
	}
	return uploadFileName, nil
}
