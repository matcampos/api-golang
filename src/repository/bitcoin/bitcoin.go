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
	stmt, err := db.Prepare("INSERT bitcoin SET quantity=?, total=?, type=?, person_id=?")
	e.CheckErr(err)

	res, err := stmt.Exec(b.Quantity, b.Total, b.Type, b.Person_id)
	e.CheckErr(err)

	id, err := res.LastInsertId()
	e.CheckErr(err)

	fmt.Println(id)

	db.Close()
	return b
}

func GetByUser() []bitcoinModel.BitCoinReportOneArray {
	db := con.Connect()
	// insert
	rows, err := db.Query("select b.total, b.quantity, b.person_id, p.name from bitcoin b inner join person p where p.id = b.person_id order by b.person_id")
	e.CheckErr(err)
	i := 0
	bitcoins := []bitcoinModel.BitCoinReportOne{}
	for rows.Next() {
		i++
		bitcoin := bitcoinModel.BitCoinReportOne{}
		err = rows.Scan(&bitcoin.Total, &bitcoin.Quantity, &bitcoin.Person_id, &bitcoin.Name)
		e.CheckErr(err)

		bitcoins = append(bitcoins, bitcoin)
	}
	bitcoinsArr := []bitcoinModel.BitCoinReportOneArray{}
	id := 0

	bitcoinsObj := bitcoinModel.BitCoinReportOneArray{}
	for i := 0; i < len(bitcoins); i++ {
		if id != 0 && id != bitcoins[i].Person_id {
			bitcoinsArr = append(bitcoinsArr, bitcoinsObj)
			bitcoinsObj = bitcoinModel.BitCoinReportOneArray{}
			bitcoinsObj.Person_id = bitcoins[i].Person_id
			bitcoinsObj.Person_name = bitcoins[i].Name
			bitcoinsObj.Bitcoins = append(bitcoinsObj.Bitcoins, bitcoins[i])
		} else if id == 0 {
			bitcoinsObj.Person_id = bitcoins[i].Person_id
			bitcoinsObj.Person_name = bitcoins[i].Name
			bitcoinsObj.Bitcoins = append(bitcoinsObj.Bitcoins, bitcoins[i])
		} else if id == bitcoins[i].Person_id {
			bitcoinsObj.Bitcoins = append(bitcoinsObj.Bitcoins, bitcoins[i])
		}
		if i == (len(bitcoins) - 1) {
			bitcoinsArr = append(bitcoinsArr, bitcoinsObj)
		}
		id = bitcoins[i].Person_id
	}

	db.Close()
	return bitcoinsArr
}

func GetByDay(day1 string, day2 string) []bitcoinModel.BitcoinDateReport {
	db := con.Connect()

	var d1 = day1

	var d2 = day2
	rows, err := db.Query("select * from bitcoin where date between ? and ?", d1, d2)
	e.CheckErr(err)
	i := 0
	bitcoins := []bitcoinModel.BitcoinDateReport{}
	for rows.Next() {
		i++
		bitcoin := bitcoinModel.BitcoinDateReport{}
		err = rows.Scan(&bitcoin.Id, &bitcoin.Quantity, &bitcoin.Total, &bitcoin.Date, &bitcoin.Type, &bitcoin.Person_id)
		e.CheckErr(err)
		bitcoins = append(bitcoins, bitcoin)
	}
	db.Close()
	return bitcoins
}
