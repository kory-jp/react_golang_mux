package usecase

import (
	"database/sql"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository interface {
	TransStore(*sql.Tx, domain.Todo) (int64, error)
	FindByUserId(int, int) (domain.Todos, float64, error)
	Erasure(int, int) error
	FindByIdAndUserId(int, int) (*domain.Todo, error)
	TransOverwrite(*sql.Tx, domain.Todo) error
	ChangeBoolean(int, int, domain.Todo) error
}
