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
	err = interactor.TodoRepository.Store(t)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		err = errors.New("保存に失敗しました")
		return
	}
	return
}

func (interactor *TodoInteractor) Todos(id int) (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindByUserId(id)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}

func (interactor *TodoInteractor) TodoByIdAndUserId(id int, userId int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindByIdAndUserId(id, userId)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}

func (interactor *TodoInteractor) Change(t domain.Todo) (mess TodoMessage, err error) {
	err = interactor.TodoRepository.Overwrite(t)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		err = errors.New("更新に失敗しました")
		return
	}
	return
}

func (interactor *TodoInteractor) Remove(id int) (mess TodoMessage, err error) {
	err = interactor.TodoRepository.Erasure(id)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		err = errors.New("削除に失敗しました")
		return
	}
	return
}
