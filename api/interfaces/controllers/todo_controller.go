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
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

type TodosError struct {
	Error string
}

func (serr *TodosError) MakeErr(mess string) (errStr string) {
	err := errors.New(mess)
	todosErr := &TodosError{Error: err.Error()}
	e, _ := json.Marshal(todosErr)
	errStr = string(e)
	return
}

type ResponseFormat struct {
	Todos   domain.Todos `json:"todos"`
	SumPage float64      `json:"sumPage"`
	Message string       `json:"message"`
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func GetUserId(r *http.Request) (userId int, err error) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	userId = session.Values["userId"].(int)
	return userId, nil
}

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {

	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("画像の容量が大きく保存できません")
		fmt.Fprintln(w, errStr)
		return
	}
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("画像の取り込み失敗しました")
			fmt.Fprintln(w, errStr)
			return
		}
	} else {
		defer file.Close()
		// 画像を保存するimgディレクトリが存在しない場合は作成する
		err = os.MkdirAll("./img", os.ModePerm)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		// サーバー側に保存するために空ファイルを作成
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		imageFilePath := "./img/" + uploadFileName
		formatPath := filepath.Clean(imageFilePath)
		saveImage, err = os.Create(formatPath)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		defer saveImage.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("保存処理に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}

	todoType := new(domain.Todo)
	todoType.UserID = userId
	todoType.Title = r.Form.Get("title")
	todoType.Content = r.Form.Get("content")
	todoType.ImagePath = uploadFileName

	mess, err := controller.Interactor.Add(*todoType)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("保存処理に失敗しました")
		fmt.Fprintln(w, errStr)
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) Index(w http.ResponseWriter, r *http.Request) {

	// URLから取得したいページ番目の情報
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	todos, sumPage, err := controller.Interactor.Todos(userId, page)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	res := ResponseFormat{
		Todos:   todos,
		SumPage: sumPage,
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonResponse))
}

func (controller *TodoController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	todo, err := controller.Interactor.TodoByIdAndUserId(id, userId)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}

	fmt.Fprintln(w, string(jsonTodo))
}

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("画像の容量が大きく保存できません")
		fmt.Fprintln(w, errStr)
		return
	}
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("画像の取り込みに失敗しました")
			fmt.Fprintln(w, errStr)
			return
		}
	} else {
		defer file.Close()
		err = os.MkdirAll("./img", os.ModePerm)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		imageFilePath := "./img/" + uploadFileName
		formatPath := filepath.Clean(imageFilePath)
		saveImage, err = os.Create(formatPath)
		if err != nil {
			//  "サーバ側でファイル確保できませんでした
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		defer saveImage.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			// アップロードしたファイルの書き込みに失敗しました。"
			fmt.Println(err)
			log.Println(err)
			errStr := new(TodosError).MakeErr("サーバーで障害が発生しました")
			fmt.Fprintln(w, errStr)
			return
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("保存処理に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	todoType := new(domain.Todo)
	todoType.ID, _ = strconv.Atoi(r.Form.Get("id"))
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

	// -------
	mess, err := controller.Interactor.UpdateTodo(*todoType)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) IsFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	todoType := new(domain.Todo)
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("保存処理に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	mess, err := controller.Interactor.IsFinishedTodo(id, *todoType, userId)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("削除処理に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	mess, err := controller.Interactor.DeleteTodo(id, userId)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) DeleteInIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("削除処理に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	todos, sumPage, mess, err := controller.Interactor.DeleteTodoInIndex(id, userId, page)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	res := ResponseFormat{
		Todos:   todos,
		SumPage: sumPage,
		Message: mess.Message,
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(TodosError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonResponse))
}
