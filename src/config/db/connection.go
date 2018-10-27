package db

import (
	"database/sql"

	e "../../helpers/error"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@/go?charset=utf8")
	e.CheckErr(err)
	return db
}
