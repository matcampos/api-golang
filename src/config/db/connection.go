package db

import (
	"database/sql"

	e "../../helpers/error"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:my-secret-pw@/bitcoin?charset=utf8")
	e.CheckErr(err)
	return db
}
