package infrastructure

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kory-jp/react_golang_mux/api/interfaces/controllers"
	"github.com/rs/cors"
)

func Init() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}

	r := mux.NewRouter()
	todoController := controllers.NewTodoController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())
	sessionController := controllers.NewSessionController(NewSqlHandler())
	r.Methods("POST").Path("/api/new").HandlerFunc(todoController.Create)
	r.Methods("GET").Path("/api/todos").HandlerFunc(todoController.Index)
	r.Methods("GET").Path("/api/todos/{id:[0-9]+}").HandlerFunc(todoController.Show)
	r.Methods("POST").Path("/api/registration").HandlerFunc(userController.Create)
	r.Methods("POST").Path("/api/login").HandlerFunc(sessionController.Login)
	r.Methods("GET").Path("/api/authenticate").HandlerFunc(sessionController.Authenticate)
	r.Methods("DELETE").Path("/api/logout").HandlerFunc(sessionController.Logout)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
