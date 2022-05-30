package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	users "github.com/kory-jp/react_golang_mux/api/interfaces/database/users"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/users"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	User    *domain.User `json:"user"`
}

func (res *Response) SetResp(status int, mess string, user *domain.User) (resStr string) {
	response := &Response{status, mess, user}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &users.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	user, err := controller.Interactor.Add(*userType)
	if err != nil {
		errStr := err.Error()
		errStr1 := strings.Replace(errStr, "Error 1062: Duplicate entry", "入力された", 1)
		errStr2 := strings.Replace(errStr1, "for key 'email'", "既に登録されています。", 1)
		resStr := new(Response).SetResp(400, errStr2, nil)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "新規登録完了しました", user)
	fmt.Fprintln(w, resStr)
}
