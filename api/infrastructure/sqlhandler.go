package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/infrastructure/mysql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kory-jp/react_golang_mux/api/config"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.UserName,
		config.Config.Password,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBname,
	)
	conn, err := sql.Open(config.Config.SQLDriver, DSN)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	errP := conn.Ping()
	if errP != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	query := mysql.Query()
	for i := 0; i < len(query); i++ {
		_, err := conn.Exec(query[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (handler *SqlHandler) DoInTx(f func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := handler.Conn.Begin()
	if err != nil {
		fmt.Println(err)
		log.Panicln(err)
		return nil, err
	}

	v, err := f(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

func (r SqlRow) Err() error {
	return r.Rows.Err()
}
