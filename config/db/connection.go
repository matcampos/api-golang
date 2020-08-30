package db

import (
	"database/sql"
	"os"

	e "../../helpers/error"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_USER")+":"+os.Getenv("DATABASE_PASSWORD")+"@/"+os.Getenv("DATABASE")+"?charset=utf8")
	e.CheckErr(err)
	return db
}
