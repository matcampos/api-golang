package db

import (
	"database/sql"
	"os"

	e "../../helpers/error"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_USER")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE")+"?charset=utf8")
	e.CheckErr(err)
	return db
}
