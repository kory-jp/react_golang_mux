package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TaskCardInteractor struct {
	TaskCardRepository TaskCardRepository
}

type TaskCardMessage struct {
	Message string
}

// --- 新規登録 ---
func (interactor *TaskCardInteractor) Add(t domain.TaskCard) (mess *TaskCardMessage, err error) {
	if err = t.TaskCardValidate(); err == nil {
		err = interactor.TaskCardRepository.Store(t)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("保存に失敗しました")
			return nil, err
		}
		mess = &TaskCardMessage{
			Message: "保存しました",
		}
		return mess, nil
	}
	return nil, err
}

// --- 一覧取得 ---
func (interactor *TaskCardInteractor) TaskCards(todoId int, userId int, page int) (taskCards domain.TaskCards, sumPage float64, err error) {
	if todoId == 0 || userId == 0 || page == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}

	taskCards, sumPage, err = interactor.TaskCardRepository.FindByTodoIdAndUserId(todoId, userId, page)
	if err != nil {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}

	return taskCards, sumPage, nil
}

// --- 詳細取得 ---
func (interactor *TaskCardInteractor) TaskCardByIdAndUserId(taskCardId int, userId int) (taskCard *domain.TaskCard, err error) {
	if taskCardId == 0 || userId == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}

	taskCard, err = interactor.TaskCardRepository.FindByIdAndUserId(taskCardId, userId)
	if err != nil {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}

	return taskCard, nil
}

// --- 新規登録 ---
func (interactor *TaskCardInteractor) UpdateTaskCard(t domain.TaskCard) (mess *TaskCardMessage, err error) {
	if err = t.TaskCardValidate(); err == nil {
		err = interactor.TaskCardRepository.Overwrite(t)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("更新に失敗しました")
			return nil, err
		}
		mess = &TaskCardMessage{
			Message: "更新しました",
		}
		return mess, nil
	}
	return nil, err
}
