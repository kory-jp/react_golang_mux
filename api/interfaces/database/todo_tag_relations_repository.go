package database

import (
	"database/sql"
	"fmt"
	"log"
)

type TodTagRelationsRepository struct {
	SqlHandler
}

var CreateTodoTagRelationsState = `
	insert into
		todo_tag_relations(
			todo_id,
			tag_id
		)
	value (?, ?)
`

func (repo *TodTagRelationsRepository) TransStore(tx *sql.Tx, todoId int64, tagIds []int) (err error) {
	for _, v := range tagIds {
		_, err = repo.TransExecute(tx, CreateTodoTagRelationsState, todoId, v)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return err
		}
	}
	return err
}
