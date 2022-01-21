package infrastructure

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kory-jp/react_golang_mux/api/interfaces/controllers"
	"github.com/rs/cors"
)

func Init() {
	r := mux.NewRouter()
	todoController := controllers.NewTodoController(NewSqlHandler())
	controllers.NewUserController(NewSqlHandler())
	r.Methods("POST").Path("/api").HandlerFunc(todoController.Create)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
