package usecase

import (
	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TaskCardRepository interface {
	Store(domain.TaskCard) error
}
