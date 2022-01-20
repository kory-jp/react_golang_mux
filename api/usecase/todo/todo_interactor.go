package usecase

import (
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(t domain.Todo) (todo domain.Todo, err error) {
	identifier, err := interactor.TodoRepository.Store(t)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		return
	}
	todo, err = interactor.TodoRepository.FindById(identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		return
	}
	return
}
