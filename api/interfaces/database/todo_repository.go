package database

import (
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t domain.Todo) (id int, err error) {
	result, err := repo.Execute(`
		insert into
			todos(
				content
			)
		value (?)
	`, t.Content)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	id64, err := result.LastInsertId()
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		return
	}
	id = int(id64)
	return
}

func (repo *TodoRepository) FindById(identifier int) (todo domain.Todo, err error) {
	row, err := repo.Query(`
		select
			id,
			content
		from
			todos
		where
			id = ?
	`, identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	defer row.Close()

	var id int
	var content string
	row.Next()
	if err = row.Scan(&id, &content); err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	todo.ID = id
	todo.Content = content
	return
}
