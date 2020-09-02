package userRepository

import (
	"fmt"

	con "api-golang/config/db"
	userModel "api-golang/models/user"

	_ "github.com/go-sql-driver/mysql"
)

func Create(u userModel.User) (userModel.User, error) {
	db, err := con.Connect()
	if err != nil {
		return u, err
	}

	stmt, err := db.Prepare("INSERT user SET name=?, email=?, password=?, birthdate=?")

	if err != nil {
		db.Close()
		return u, err
	}

	res, err := stmt.Exec(u.Name, u.Email, u.Password, u.Birthdate)

	if err != nil {
		db.Close()
		return u, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		db.Close()
		return u, err
	}

	fmt.Println(id)

	u.Id = int(id)

	db.Close()
	return u, nil
}
