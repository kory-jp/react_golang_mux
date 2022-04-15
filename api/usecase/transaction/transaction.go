package transaction

import (
	"database/sql"
)

type SqlHandler interface {
	DoInTx(func(tx *sql.Tx) (interface{}, error)) (interface{}, error)
}
