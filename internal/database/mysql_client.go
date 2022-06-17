package database

import (
	"database/sql"
	"fmt"
)

type MySqlClient struct {
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("Cannot create db tentat: %s", err.Error())
		panic("..")
	}

	return db
}
