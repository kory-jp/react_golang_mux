package main

import (
	"fmt"

	"github.com/kory-jp/react_golang_mux/api/infrastructure"
	"github.com/kory-jp/react_golang_mux/api/infrastructure/seeds/development/seed"
)

func main() {
	con := infrastructure.NewSqlHandler()
	if err := seed.UsersSeed(con); err != nil {
		fmt.Println(err)
		return
	}
	if err := seed.TodosSeed(con); err != nil {
		fmt.Println(err)
		return
	}
	if err := seed.TagsSeed(con); err != nil {
		fmt.Println(err)
		return
	}

	if err := seed.TodoTagRelationsSeed(con); err != nil {
		fmt.Println(err)
		return
	}

	if err := seed.TaskCardsSeed(con); err != nil {
		fmt.Println(err)
		return
	}
}
