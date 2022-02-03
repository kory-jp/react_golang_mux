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
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/user"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type UserValidError struct {
	Detail string
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
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	user, err := controller.Interactor.Add(*userType)
	if err != nil {
		errStr := err.Error()
		errStr1 := strings.Replace(errStr, "Error 1062: Duplicate entry", "入力された", 1)
		errStr2 := strings.Replace(errStr1, "for key 'email'", "既に登録されています。", 1)
		fmt.Println(errStr2)
		validErr := &UserValidError{Detail: errStr2}
		e, _ := json.Marshal(validErr)
		fmt.Fprintln(w, string(e))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	fmt.Fprintln(w, string(jsonUser))
}
