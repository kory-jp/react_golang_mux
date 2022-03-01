package seed

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kory-jp/react_golang_mux/api/config"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.Config.UserName,
		config.Config.Password,
		config.Config.DBPort,
		config.Config.DBname,
	)
	conn, err := sql.Open(config.Config.SQLDriver, DSN)
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗", err)
	} else {
		fmt.Println("データベース接続成功")
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
