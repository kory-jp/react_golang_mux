package usecase

import "github.com/kory-jp/react_golang_mux/api/domain"

type SessionRepository interface {
	FindByEmail(domain.User) (*domain.User, error)
	FindById(int) (*domain.User, error)
}
