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

// --- Todo新規追加 ---
func (interactor *TodoInteractor) Add(t domain.Todo, tagIds []int) (mess *TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		_, err = interactor.Transaction.DoInTx(func(tx *sql.Tx) (interface{}, error) {
			todoId, err := interactor.TodoRepository.TransStore(tx, t)
			if err != nil {
				return nil, err
			}
			if len(tagIds) != 0 {
				err = interactor.TodoTagRelationsRepository.TransStore(tx, todoId, tagIds)
			}
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
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

// --- Todo一覧取得 ---
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

// --- Todo詳細情報取得 ---
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

// --- タグ検索 ---
func (interactor *TodoInteractor) SearchTag(tagId int, userId int, page int) (todos domain.Todos, sumPage float64, err error) {
	if tagId == 0 || userId == 0 || page == 0 {
		err = errors.New("データ取得に失敗しました")
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}

	todos, sumPage, err = interactor.TodoRepository.FindByTagId(tagId, userId, page)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("データ取得に失敗しました")
		return nil, 0, err
	}

	return todos, sumPage, nil
}

// --- Todo更新 ---
func (interactor *TodoInteractor) UpdateTodo(t domain.Todo, tagIds []int) (mess *TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		_, err = interactor.Transaction.DoInTx(func(tx *sql.Tx) (interface{}, error) {
			err = interactor.TodoRepository.TransOverwrite(tx, t)
			if err != nil {
				return nil, err
			}
			err = interactor.TodoTagRelationsRepository.TransOverwrite(tx, t.ID, tagIds)
			fmt.Println(err)
			return nil, err
		})
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

// --- Todo完了未完了変更 ---
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

// --- Todo削除 ---
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

// --- Todo削除 + 一覧取得
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
