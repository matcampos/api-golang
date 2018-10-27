package bitcoinRepository

import (
	"fmt"

	con "../../config/db"
	e "../../helpers/error"
	"../../models/bitcoin"
	_ "github.com/go-sql-driver/mysql"
)

func Create(b bitcoinModel.Bitcoin) bitcoinModel.Bitcoin {
	db := con.Connect()
	// insert
	stmt, err := db.Prepare("INSERT bitcoin SET quantity=?, total=?")
	e.CheckErr(err)

	res, err := stmt.Exec(b.Quantity, b.Total)
	e.CheckErr(err)

	id, err := res.LastInsertId()
	e.CheckErr(err)

	fmt.Println(id)

	db.Close()
	return b
}
