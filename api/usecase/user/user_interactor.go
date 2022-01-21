package usecase

import (
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	user, err = interactor.UserRepository.FindById(identifier)
	if err != nil {
		return
	}
	return
}
