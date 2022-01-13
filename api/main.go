package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kory-jp/react_golang_mux/api/config"
)

func main() {
	// controller
	http.HandleFunc("/", echoHello)
	// port
	http.ListenAndServe(":8000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hello World</h1>")
	fmt.Fprintf(w, config.Config.Port)
	log.Println("test")
}
