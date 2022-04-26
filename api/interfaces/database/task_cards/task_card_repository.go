package database

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database/task_cards/mysql"
)

type TaskCardRepository struct {
	database.SqlHandler
}

var allTodosCount float64

func (repo *TaskCardRepository) Store(t domain.TaskCard) (err error) {
	_, err = repo.Execute(mysql.CreateTaskCardState, t.UserID, t.TodoID, t.Title, t.Purpose, t.Content, t.Memo, false, time.Now())
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return nil
}

func (repo *TaskCardRepository) FindByTodoIdAndUserId(todoId int, userId int, page int) (taskCards domain.TaskCards, sumPage float64, err error) {
	row, err := repo.Query(mysql.SumTaskCardItemsState, userId, todoId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&allTodosCount)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
	}

	err = row.Err()
	if err != nil {
		fmt.Println(err)
	}
	row.Close()

	sumPage = math.Ceil(allTodosCount / 5)

	var offsetNum int
	if page == 1 {
		offsetNum = 0
	} else {
		offsetNum = (page - 1) * 5
	}
	rows, err := repo.Query(mysql.GetTaskCardsState, userId, todoId, offsetNum)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var taskCard domain.TaskCard
		err = rows.Scan(
			&taskCard.ID,
			&taskCard.UserID,
			&taskCard.TodoID,
			&taskCard.Title,
			&taskCard.Purpose,
			&taskCard.Content,
			&taskCard.Memo,
			&taskCard.IsFinished,
			&taskCard.CreatedAt,
		)
		if err != nil {
			fmt.Println("ID,UserIDと一致するTodoが存在していない")
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
		taskCards = append(taskCards, taskCard)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	row.Close()
	return taskCards, sumPage, err
}
