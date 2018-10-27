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
	stmt, err := db.Prepare("INSERT person SET name=?, age=?")
	e.CheckErr(err)

	res, err := stmt.Exec(u.Name, u.Age)
	e.CheckErr(err)

	id, err := res.LastInsertId()
	e.CheckErr(err)

	fmt.Println(id)

	db.Close()
	return u
}

func GetAll() []userModel.User {
	db := con.Connect()
	// insert
	rows, err := db.Query("SELECT * FROM person")
	e.CheckErr(err)
	i := 0
	users := []userModel.User{}
	for rows.Next() {
		i++
		user := userModel.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		e.CheckErr(err)

		println(user.Name)
		users = append(users, user)
	}
	db.Close()
	return users
}
