package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TodoTagRelationsRepository struct {
	SqlHandler
}

var CreateTodoTagRelationsState = `
	insert into
		todo_tag_relations(
			todo_id,
			tag_id,
			created_at
		)
	value (?, ?)
`

var DeleteTodoTagRelationsState = `
		delete from
			todo_tag_relations
		where
			todo_tag_relations.todo_id = ?
`

func (repo *TodoTagRelationsRepository) TransStore(tx *sql.Tx, todoId int64, tagIds []int) (err error) {
	for _, v := range tagIds {
		_, err = repo.TransExecute(tx, CreateTodoTagRelationsState, todoId, v, time.Now())
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return err
		}
	}
	return err
}

func (repo *TodoTagRelationsRepository) TransOverwrite(tx *sql.Tx, todoId int, tagIds []int) (err error) {
	_, err = repo.TransExecute(tx, DeleteTodoTagRelationsState, todoId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
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
