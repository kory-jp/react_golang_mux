package infrastructure

import (
	"log"
	"net/http"

	sessions "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
	tags "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/tags"
	taskCards "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/task_cards"
	todos "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/todos"
	users "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/users"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func Init() {
	log.SetFlags(log.Ltime | log.Llongfile)
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Println(err)
		log.Panicln(err)
	}

	r := mux.NewRouter()
	todoController := todos.NewTodoController(NewSqlHandler())
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
	// ----- 画像配信URL ---------
	r.PathPrefix("/api/img/").Handler(http.StripPrefix("/api/img/", http.FileServer(http.Dir("./assets/dev/img"))))
	// -----
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
