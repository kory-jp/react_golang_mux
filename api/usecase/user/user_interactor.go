package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user *domain.User, err error) {
	if err = u.UserValidate(); err == nil {
		u.Password = u.Encrypt(u.Password)
		identifier, err := interactor.UserRepository.Store(u)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("データ保存に失敗しました")
			return nil, err
		} else {
			user, err = interactor.UserRepository.FindById(identifier)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				err = errors.New("データ取得に失敗しました")
				return nil, err
			}
			return user, nil
		}
	}
	return nil, err
}
