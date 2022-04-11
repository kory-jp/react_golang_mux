package transaction

import (
	"database/sql"
)

type SqlHandler interface {
	DoInTx(f func(tx *sql.Tx) (interface{}, error)) (interface{}, error)
}
