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
