package apiBitcoin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Bitcoin struct {
	BPI BPI `json:"bpi"`
}

type BPI struct {
	USD CoutryPrice `json:"USD"`
	BRL CoutryPrice `json:"BRL"`
}
type CoutryPrice struct {
	Code        string  `json:"code"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	Rate_float  float64 `json:"rate_float"`
}

func DecodeData() (float64, error) {

	response, err := http.Get(os.Getenv("BITCOIN_API_ENDPOINT"))
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 0, err
	}

	var bit Bitcoin
	json.Unmarshal([]byte(contents), &bit)
	return bit.BPI.BRL.Rate_float, nil
}
