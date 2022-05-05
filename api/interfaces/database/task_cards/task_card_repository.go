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

// -- 新規作成 ---
// ---
func (repo *TaskCardRepository) Store(t domain.TaskCard) (err error) {
	_, err = repo.Execute(mysql.CreateTaskCardState, t.UserID, t.TodoID, t.Title, t.Purpose, t.Content, t.Memo, false)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return nil
}

// --- 一覧取得 ---
// ---
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

// --- 詳細取得 ---
// ---
func (repo *TaskCardRepository) FindByIdAndUserId(taskCardId int, userId int) (taskCard *domain.TaskCard, err error) {
	row, err := repo.Query(mysql.ShowTaskCardState, taskCardId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	defer row.Close()

	var (
		id         int
		user_id    int
		todo_id    int
		title      string
		purpose    string
		content    string
		memo       string
		isFinished bool
		created_at time.Time
	)

	row.Next()
	if err = row.Scan(
		&id,
		&user_id,
		&todo_id,
		&title,
		&purpose,
		&content,
		&memo,
		&isFinished,
		&created_at,
	); err != nil {
		fmt.Println("ID,UserIDと一致するTodoが存在していない")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	row.Close()
	taskCard = &domain.TaskCard{
		ID:         id,
		UserID:     user_id,
		TodoID:     todo_id,
		Title:      title,
		Purpose:    purpose,
		Content:    content,
		Memo:       memo,
		IsFinished: isFinished,
		CreatedAt:  created_at,
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
	}
	row.Close()
	return taskCard, nil
}

func (repo *TaskCardRepository) Overwrite(t domain.TaskCard) (err error) {
	_, err = repo.Execute(mysql.UpdateTaskCardState, t.Title, t.Purpose, t.Content, t.Memo, t.ID, t.UserID)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return nil
}

func (repo *TaskCardRepository) ChangeBoolean(id int, userId int, taskCard domain.TaskCard) (err error) {
	_, err = repo.Execute(mysql.ChangeBoolState, taskCard.IsFinished, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

func (repo *TaskCardRepository) Erasure(taskCardId int, userId int) (err error) {
	_, err = repo.Execute(mysql.DeleteTaskCardState, taskCardId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

func (repo *TaskCardRepository) GetCounts(todoId int, userId int) (incompleteTaskCount int, err error) {
	row, err := repo.Query(mysql.GetIncompleteTaskCount, todoId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}

	defer row.Close()
	for row.Next() {
		err = row.Scan(&incompleteTaskCount)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return 0, err
		}
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	row.Close()

	return incompleteTaskCount, nil
}
