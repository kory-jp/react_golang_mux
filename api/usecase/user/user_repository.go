package usecase

import "github.com/kory-jp/react_golang_mux/api/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (*domain.User, error)
}
