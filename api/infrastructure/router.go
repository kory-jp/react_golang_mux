package infrastructure

import (
	"log"
	"net/http"
	"os"
	"strings"

	tags "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/tags"

	"github.com/kory-jp/react_golang_mux/api/infrastructure/aws"

	sessions "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
	taskCards "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/task_cards"
	todos "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/todos"
	users "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/users"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	log.SetFlags(log.Ltime | log.Llongfile)

	r := mux.NewRouter()
	todoController := todos.NewTodoController(NewSqlHandler(), aws.NewS3())
	userController := users.NewUserController(NewSqlHandler())
	sessionController := sessions.NewSessionController(NewSqlHandler())
	tagController := tags.NewTagController(NewSqlHandler())
	taskCardController := taskCards.NewTaskCardController(NewSqlHandler())
	r.Methods("POST").Path("/api/registration").HandlerFunc(userController.Create)
	r.Methods("POST").Path("/api/login").HandlerFunc(sessionController.Login)
	r.Methods("GET").Path("/api/authenticate").HandlerFunc(sessionController.Authenticate)
	r.Methods("DELETE").Path("/api/logout").HandlerFunc(sessionController.Logout)
	r.Methods("POST").Path("/api/new").HandlerFunc(todoController.Create)
	r.Methods("GET").Path("/api/todos").HandlerFunc(todoController.Index)
	r.Methods("GET").Path("/api/todos/{id:[0-9]+}").HandlerFunc(todoController.Show)
	r.Methods("GET").Path("/api/todos/search").HandlerFunc(todoController.Search)
	r.Methods("POST").Path("/api/todos/update/{id:[0-9]+}").HandlerFunc(todoController.Update)
	r.Methods("POST").Path("/api/todos/isfinished/{id:[0-9]+}").HandlerFunc(todoController.IsFinished)
	r.Methods("DELETE").Path("/api/todos/delete/{id:[0-9]+}").HandlerFunc(todoController.Delete)
	r.Methods("DELETE").Path("/api/todos/deleteinindex/{id:[0-9]+}").HandlerFunc(todoController.DeleteInIndex)
	r.Methods("GET").Path("/api/tag").HandlerFunc(tagController.Index)
	r.Methods("POST").Path("/api/taskcard/new").HandlerFunc(taskCardController.Create)
	r.Methods("GET").Path("/api/todo/{id:[0-9]+}/taskcard").HandlerFunc(taskCardController.Index)
	r.Methods("GET").Path("/api/taskcard/{id:[0-9]+}").HandlerFunc(taskCardController.Show)
	r.Methods("POST").Path("/api/taskcard/{id:[0-9]+}").HandlerFunc(taskCardController.Update)
	r.Methods("POST").Path("/api/taskcard/isfinished/{id:[0-9]+}").HandlerFunc(taskCardController.IsFinished)
	r.Methods("DELETE").Path("/api/taskcard/{id:[0-9]+}").HandlerFunc(taskCardController.Delete)
	r.Methods("GET").Path("/api/taskcard/incompletetaskcount/{id:[0-9]+}").HandlerFunc(taskCardController.IncompleteTaskCount)
	// ----- 画像配信URL ---------
	r.PathPrefix("/api/img/").Handler(http.StripPrefix("/api/img/", http.FileServer(http.Dir("./assets/dev/img"))))
	// -----
	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("ALLOWED_ORIGINS"), " "),
		AllowCredentials: true,
		AllowedMethods:   strings.Split(os.Getenv("ALLOWED_METHODS"), " "),
		AllowedHeaders:   strings.Split(os.Getenv("ALLOWED_HEADERS"), " "),
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
