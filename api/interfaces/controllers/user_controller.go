package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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
		validErr := &UserValidError{Detail: err.Error()}
		e, _ := json.Marshal(validErr)
		fmt.Fprintln(w, string(e))
		return
	}
	fmt.Println(user)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	fmt.Fprintln(w, string(jsonUser))
}
