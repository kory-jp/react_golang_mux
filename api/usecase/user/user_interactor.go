package usecase

import (
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	if err = domain.UserValidate(&u); err != nil {
		log.Println(err)
		fmt.Println(err)
	} else {
		identifier, err := interactor.UserRepository.Store(u)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Println(err)
		}
		user, err = interactor.UserRepository.FindById(identifier)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Println(err)
		}
	}
	return
}
