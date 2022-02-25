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

type ResponseFormat struct {
	Todos   domain.Todos `json:"todos"`
	SumPage float64      `json:"sumPage"`
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

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {

	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		fmt.Println("No Image File", err)
		// fmt.Fprintln(w, "ファイルアップロードを確認できませんでした。")
	} else {
		// 画像を保存するimgディレクトリが存在しない場合は作成する
		err = os.MkdirAll("./img", os.ModePerm)
		if err != nil {
			fmt.Fprintln(w, "サーバーで障害が発生しました")
			return
		}
		// サーバー側に保存するために空ファイルを作成
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		saveImage, err = os.Create("./img/" + uploadFileName)
		if err != nil {
			fmt.Fprintln(w, "サーバ側でファイル確保できませんでした。")
			return
		}
		defer saveImage.Close()
		defer file.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			fmt.Println(err)
			fmt.Println("アップロードしたファイルの書き込みに失敗しました。")
			os.Exit(1)
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	todoType := new(domain.Todo)
	todoType.UserID, _ = strconv.Atoi(r.Form.Get("user_id"))
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
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) Index(w http.ResponseWriter, r *http.Request) {

	// URLから取得したいページ番目の情報
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	session, _ := store.Get(r, "session")
	todos, sumPage, err := controller.Interactor.Todos(session.Values["userId"].(int), page)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		err := errors.New("データ取得に失敗しました")
		todosErr := &TodosError{Error: err.Error()}
		e, _ := json.Marshal(todosErr)
		fmt.Fprintln(w, string(e))
	}
	res := ResponseFormat{
		Todos:   todos,
		SumPage: sumPage,
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonResponse))
}

func (controller *TodoController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	session, err := store.Get(r, "session")
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	todo, err := controller.Interactor.TodoByIdAndUserId(id, session.Values["userId"].(int))
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		err := errors.New("データ取得に失敗しました")
		todosErr := &TodosError{Error: err.Error()}
		e, _ := json.Marshal(todosErr)
		fmt.Fprintln(w, string(e))
	}

	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(jsonTodo))
}

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadFileName string
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		fmt.Println("No Image File by Edit", err)
	} else {
		err = os.MkdirAll("./img", os.ModePerm)
		if err != nil {
			fmt.Fprintln(w, "サーバーで障害が発生しました")
			return
		}
		var saveImage *os.File
		uploadFileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
		saveImage, err = os.Create("./img/" + uploadFileName)
		if err != nil {
			fmt.Fprintln(w, "サーバ側でファイル確保できませんでした。")
			return
		}
		defer saveImage.Close()
		defer file.Close()
		size, err := io.Copy(saveImage, file)
		if err != nil {
			fmt.Println(err)
			fmt.Println("アップロードしたファイルの書き込みに失敗しました。")
			os.Exit(1)
		}
		fmt.Println("書き込んだByte数=>", size)
	}

	todoType := new(domain.Todo)
	todoType.ID, _ = strconv.Atoi(r.Form.Get("id"))
	todoType.UserID, _ = strconv.Atoi(r.Form.Get("user_id"))
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
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) IsFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	todoType := new(domain.Todo)
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}

	fmt.Println(todoType)

	session, err := store.Get(r, "session")
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	mess, err := controller.Interactor.IsFinishedTodo(id, *todoType, session.Values["userId"].(int))
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonMess))
}

func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}

	mess, err := controller.Interactor.DeleteTodo(id)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	jsonMess, err := json.Marshal(mess)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonMess))
}
