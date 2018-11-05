package apiBitcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Bitcoin struct {
	Data Data `json:"data"`
}

type Data struct {
	Quotes Quotes `json:"quotes"`
}
type Quotes struct {
	BRL Price `json:"BRL"`
}

type Price struct {
	Price float64 `json:"price"`
}

func DecodeData() float64 {

	response, err := http.Get("https://api.coinmarketcap.com/v2/ticker/1/?convert=BRL")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
	var bit Bitcoin
	json.Unmarshal([]byte(contents), &bit)
	println(bit.Data.Quotes.BRL.Price)
	return bit.Data.Quotes.BRL.Price
}
