package usecase

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type SessionInteractor struct {
	SessionRepository SessionRepository
}

type SessionValidError struct {
	Detail string
}

func (interactor *SessionInteractor) Login(u domain.User) (user domain.User, err error) {
	userFindByEmail, err := interactor.SessionRepository.FindByEmail(u)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		err = errors.New("認証に失敗しました")
	} else {
		// if userFindByEmail.Password == u.Encrypt(u.Password) {
		// 	user = userFindByEmail
		// } else {
		// 	err = errors.New("認証に失敗しました")
		// }
		err = bcrypt.CompareHashAndPassword([]byte(userFindByEmail.Password), []byte(u.Password))
		if err == nil {
			user = userFindByEmail
		} else {
			err = errors.New("認証に失敗しました")
		}
	}
	return
}

func (interactor *SessionInteractor) IsLoggedin(uid int) (user domain.User, err error) {
	user, err = interactor.SessionRepository.FindById(uid)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}
