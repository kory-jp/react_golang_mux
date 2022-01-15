package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kory-jp/react_golang_mux/api/config"
	"github.com/rs/cors"
)

func echoHello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hello World</h1>")
	fmt.Fprintf(w, config.Config.Port)
}

func main() {
	r := mux.NewRouter()
	r.Methods("POST").Path("/api").HandlerFunc(echoHello)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
