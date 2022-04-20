package usecase

import (
	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TagRepository interface {
	FindAll() (domain.Tags, error)
}
