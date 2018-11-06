package userRepository

import (
	"fmt"

	con "../../config/db"
	e "../../helpers/error"
	"../../models/user"
	_ "github.com/go-sql-driver/mysql"
)

func Create(u userModel.User) userModel.User {
	db := con.Connect()
	// insert
	stmt, err := db.Prepare("INSERT user SET name=?, email=?, password=?, birthdate=?")
	e.CheckErr(err)

	res, err := stmt.Exec(u.Name, u.Email, u.Password, u.Birthdate)
	e.CheckErr(err)

	id, err := res.LastInsertId()
	e.CheckErr(err)

	fmt.Println(id)

	db.Close()
	return u
}
