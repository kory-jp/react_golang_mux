package usecase

import (
	"database/sql"
)

type TodoTagRelationsRepository interface {
	TransStore(*sql.Tx, int64, []int) error
	TransOverwrite(*sql.Tx, int, []int) error
}
