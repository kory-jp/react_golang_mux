package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/usecase/transaction"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoInteractor struct {
	TodoRepository             TodoRepository
	TodoTagRelationsRepository TodoTagRelationsRepository
	Transaction                transaction.SqlHandler
}

type TodoMessage struct {
	Message string
}

func (interactor *TodoInteractor) Add(t domain.Todo, tagIds []int) (mess *TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		_, err = interactor.Transaction.DoInTx(func(tx *sql.Tx) (interface{}, error) {
			todoId, err := interactor.TodoRepository.Store(t)
			if err != nil {
				return nil, err
			}
			err = interactor.TodoTagRelationsRepository.Store(todoId, tagIds)
			fmt.Println("////IsERR?////", err)
			return nil, err
		})
		if err != nil {
			fmt.Println("$$$$$IsERR$$$$", err)
			log.Println(err)
			err = errors.New("保存に失敗しました")
			return nil, err
		}
		mess = &TodoMessage{
			Message: "保存しました",
		}
		return mess, nil
	}
	return nil, err
}

func (interactor *TodoInteractor) Todos(userId int, page int) (todos domain.Todos, sumPage float64, err error) {
	if userId == 0 || page == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}

	todos, sumPage, err = interactor.TodoRepository.FindByUserId(userId, page)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("データ取得に失敗しました")
		return nil, 0, err
	}
	return todos, sumPage, nil
}

func (interactor *TodoInteractor) TodoByIdAndUserId(id int, userId int) (todo *domain.Todo, err error) {
	if id == 0 || userId == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	todo, err = interactor.TodoRepository.FindByIdAndUserId(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("データ取得に失敗しました")
		return nil, err
	}
	return todo, nil
}

func (interactor *TodoInteractor) UpdateTodo(t domain.Todo) (mess *TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		err = interactor.TodoRepository.Overwrite(t)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("更新に失敗しました")
			return nil, err
		}
		mess = &TodoMessage{
			Message: "更新しました",
		}
		return mess, nil
	}
	return nil, err
}

func (interactor *TodoInteractor) IsFinishedTodo(id int, t domain.Todo, userId int) (mess *TodoMessage, err error) {
	if id == 0 || userId == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}

	err = interactor.TodoRepository.ChangeBoolean(id, userId, t)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("更新に失敗しました")
		return nil, err
	}

	todo, err := interactor.TodoRepository.FindByIdAndUserId(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("情報の取得に失敗しました")
		return nil, err
	}

	if todo.IsFinished {
		mess = &TodoMessage{
			Message: "完了しました",
		}
	} else {
		mess = &TodoMessage{
			Message: "未完了の項目が追加されました",
		}
	}
	return mess, nil
}

func (interactor *TodoInteractor) DeleteTodo(id int, userId int) (mess *TodoMessage, err error) {
	if id == 0 || userId == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}

	err = interactor.TodoRepository.Erasure(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("削除に失敗しました")
		return nil, err
	}
	mess = &TodoMessage{
		Message: "削除しました",
	}
	return mess, nil
}

func (interactor *TodoInteractor) DeleteTodoInIndex(id int, userId int, page int) (todos domain.Todos, sumPage float64, mess *TodoMessage, err error) {
	if id == 0 || userId == 0 || page == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, 0, nil, err
	}

	err = interactor.TodoRepository.Erasure(id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("削除に失敗しました")
		return nil, 0, nil, err
	}
	todos, sumPage, err = interactor.TodoRepository.FindByUserId(userId, page)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0, nil, err
	}
	mess = &TodoMessage{
		Message: "削除しました",
	}
	return todos, sumPage, mess, nil
}
