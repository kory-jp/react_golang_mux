package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/interfaces/database/todo_tag_relations/mysql"

	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
)

type TodoTagRelationsRepository struct {
	database.SqlHandler
}

func (repo *TodoTagRelationsRepository) TransStore(tx *sql.Tx, todoId int64, tagIds []int) (err error) {
	for _, v := range tagIds {
		_, err = repo.TransExecute(tx, mysql.CreateTodoTagRelationsState, todoId, v, time.Now())
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return err
		}
	}
	return err
}

func (repo *TodoTagRelationsRepository) TransOverwrite(tx *sql.Tx, todoId int, tagIds []int) (err error) {
	_, err = repo.TransExecute(tx, mysql.DeleteTodoTagRelationsState, todoId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	for _, v := range tagIds {
		_, err = repo.TransExecute(tx, mysql.CreateTodoTagRelationsState, todoId, v, time.Now())
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return err
		}
	}
	return err
}
