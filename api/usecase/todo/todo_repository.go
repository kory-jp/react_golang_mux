package usecase

import "github.com/kory-jp/react_golang_mux/api/domain"

type TodoRepository interface {
	Store(domain.Todo) (int64, error)
	FindByUserId(int, int) (domain.Todos, float64, error)
	FindByIdAndUserId(int, int) (*domain.Todo, error)
	Overwrite(domain.Todo) error
	ChangeBoolean(int, int, domain.Todo) error
	Erasure(int, int) error
}
