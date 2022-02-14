package usecase

import "github.com/kory-jp/react_golang_mux/api/domain"

type TodoRepository interface {
	Store(domain.Todo) error
	FindByUserId(int) (domain.Todos, error)
}
