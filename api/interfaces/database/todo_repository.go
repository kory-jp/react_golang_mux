package database

import (
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t domain.Todo) (err error) {
	_, err = repo.Execute(`
		insert into
			todos(
				user_id,
				title,
				content,
				image_path,
				isFinished,
				created_at
			)
		value (?, ?, ?, ?, ?, ?)
	`, t.UserID, t.Title, t.Content, t.ImagePath, false, time.Now())
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	return
}

func (repo *TodoRepository) FindByUserId(identifier int) (todos domain.Todos, err error) {
	rows, err := repo.Query(`
		select
			*
		from
			todos
		where
			user_id = ?
	`, identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		// var todosType domain.Todos
		err = rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Content,
			&todo.ImagePath,
			&todo.IsFinished,
			&todo.CreatedAt,
		)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Panicln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}
