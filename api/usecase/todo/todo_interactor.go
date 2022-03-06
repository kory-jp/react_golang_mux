package usecase

import (
	"errors"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

type TodoMessage struct {
	Message string
}

func (interactor *TodoInteractor) Add(t domain.Todo) (mess TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		err = interactor.TodoRepository.Store(t)
		if err != nil {
			log.Panicln(err)
			err = errors.New("保存に失敗しました")
			return
		}
	}
	mess.Message = "保存しました"
	return
}

func (interactor *TodoInteractor) Todos(userId int, page int) (todos domain.Todos, sumPage float64, err error) {
	todos, sumPage, err = interactor.TodoRepository.FindByUserId(userId, page)
	if err != nil {
		log.Println(err)
	}
	return
}

func (interactor *TodoInteractor) TodoByIdAndUserId(id int, userId int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindByIdAndUserId(id, userId)
	if err != nil {
		log.Println(err)
	}
	return
}

func (interactor *TodoInteractor) UpdateTodo(t domain.Todo) (mess TodoMessage, err error) {
	if err = t.TodoValidate(); err == nil {
		err = interactor.TodoRepository.Overwrite(t)
		if err != nil {
			log.Panicln(err)
			err = errors.New("更新に失敗しました")
			return
		}
	}
	mess.Message = "更新しました"
	return
}

func (interactor *TodoInteractor) IsFinishedTodo(id int, t domain.Todo, userId int) (mess TodoMessage, err error) {
	err = interactor.TodoRepository.ChangeBoolean(id, t)
	if err != nil {
		log.Panicln(err)
		err = errors.New("更新に失敗しました")
		return
	}

	todo, err := interactor.TodoRepository.FindByIdAndUserId(id, userId)
	if err != nil {
		log.Println(err)
	}

	if todo.IsFinished {
		mess.Message = "完了しました"
	} else {
		mess.Message = "未完了の項目が追加されました"
	}
	return
}

func (interactor *TodoInteractor) DeleteTodo(id int, userId int) (mess TodoMessage, err error) {
	err = interactor.TodoRepository.Erasure(id, userId)
	if err != nil {
		log.Panicln(err)
		err = errors.New("削除に失敗しました")
		return
	}
	mess.Message = "削除しました"
	return
}

func (interactor *TodoInteractor) DeleteTodoInIndex(id int, userId int, page int) (todos domain.Todos, sumPage float64, mess TodoMessage, err error) {
	err = interactor.TodoRepository.Erasure(id, userId)
	if err != nil {
		log.Panicln(err)
		err = errors.New("削除に失敗しました")
		return
	}
	todos, sumPage, err = interactor.TodoRepository.FindByUserId(userId, page)
	if err != nil {
		log.Println(err)
	}
	mess.Message = "削除しました"
	return
}
