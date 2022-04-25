package database

import (
	"fmt"
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database/task_cards/mysql"
)

type TaskCardRepository struct {
	database.SqlHandler
}

func (repo *TaskCardRepository) Store(t domain.TaskCard) (err error) {
	_, err = repo.Execute(mysql.CreateTaskCardState, t.UserID, t.TodoID, t.Title, t.Purpose, t.Content, t.Memo, false, time.Now())
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return nil
}
