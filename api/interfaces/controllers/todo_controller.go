package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
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
	// POSTされたファイルデータを取得する
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		fmt.Println("No Image File")
		fmt.Fprintln(w, "ファイルアップロードを確認できませんでした。")
	} else {
		fmt.Println("Ok! hello")
		fmt.Printf("%T", file)
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
	fmt.Println(*todoType)

	mess, err := controller.Interactor.Add(*todoType)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, mess)
}
