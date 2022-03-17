package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/user"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type UserValidError struct {
	Error string
}

func (serr *UserValidError) MakeErr(mess string) (errStr string) {
	err := errors.New(mess)
	usersErr := &UserValidError{Error: err.Error()}
	e, _ := json.Marshal(usersErr)
	errStr = string(e)
	return
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(UserValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(UserValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	user, err := controller.Interactor.Add(*userType)
	if err != nil {
		errStr := err.Error()
		errStr1 := strings.Replace(errStr, "Error 1062: Duplicate entry", "入力された", 1)
		errStr2 := strings.Replace(errStr1, "for key 'email'", "既に登録されています。", 1)
		validErr := &UserValidError{Error: errStr2}
		e, _ := json.Marshal(validErr)
		fmt.Fprintln(w, string(e))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(UserValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	fmt.Fprintln(w, string(jsonUser))
}
