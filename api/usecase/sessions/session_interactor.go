package usecase

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) Login(u domain.User) (user *domain.User, err error) {
	if u.Email == "" || u.Password == "" {
		err = errors.New("認証に失敗しました")
		return nil, err
	}
	userFindByEmail, err := interactor.SessionRepository.FindByEmail(u)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("認証に失敗しました")
		return nil, err
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(userFindByEmail.Password), []byte(u.Password))
		if err == nil {
			user = userFindByEmail
		} else {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("認証に失敗しました")
			return nil, err
		}
	}
	return user, nil
}

func (interactor *SessionInteractor) IsLoggedin(uid int) (user *domain.User, err error) {
	if uid == 0 {
		err = errors.New("認証に失敗しました")
		return nil, err
	}
	user, err = interactor.SessionRepository.FindById(uid)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("認証に失敗しました")
		return nil, err
	}
	return user, nil
}
