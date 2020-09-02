package db

import (
	"database/sql"
	"os"
)

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", os.Getenv("DATABASE_USER")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE")+"?charset=utf8")
}
