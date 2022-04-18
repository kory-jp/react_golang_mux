package database

import "database/sql"

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
	DoInTx(func(tx *sql.Tx) (interface{}, error)) (interface{}, error)
	TransExecute(*sql.Tx, string, ...interface{}) (Result, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}
