package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags((log.Llongfile))
		log.Panicln(err)
	}
	todoType := new(domain.Todo)
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		return
	}
	todo, err := controller.Interactor.Add(*todoType)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	fmt.Fprintln(w, string(jsonTodo))
}
