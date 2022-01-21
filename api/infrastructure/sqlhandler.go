package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kory-jp/react_golang_mux/api/config"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

const (
	tableNameTodo = "todos"
)

func NewSqlHandler() *SqlHandler {
	DSN := fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		config.Config.UserName,
		config.Config.Password,
		config.Config.DBPort,
		config.Config.DBname,
	)
	conn, _ := sql.Open(config.Config.SQLDriver, DSN)
	errP := conn.Ping()
	if errP != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

	cmdT := fmt.Sprintf(`
		create table if not exists %s (
			id integer primary key auto_increment,
			content text
		)`, tableNameTodo)
	_, errT := conn.Exec(cmdT)
	if errT != nil {
		fmt.Println("error!2")
		log.SetFlags(log.Llongfile)
		log.Println(errT)
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
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
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
