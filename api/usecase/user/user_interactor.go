package usecase

import (
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	if err = u.UserValidate(); err == nil {
		u.Password = u.Encrypt(u.Password)
		identifier, validErr := interactor.UserRepository.Store(u)
		err = validErr
		if err != nil {
			log.Println(err)
			return
		} else {
			user, err = interactor.UserRepository.FindById(identifier)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return
}
