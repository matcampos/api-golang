package bitcoinController

import (
	"encoding/json"
	"net/http"

	jsonResponse "../../helpers/json-response"
	"../../models/apiBitcoin"
	bitcoinModel "../../models/bitcoin"
	success "../../models/success"
	bitcoinRepository "../../repository/bitcoin"
)

func CreatePurchase(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	b := bitcoinModel.Bitcoin{}
	err := decoder.Decode(&b)

	if err != nil {
		panic(err)
	}
	b.Total = float64(b.Quantity) * apiBitcoin.DecodeData()
	b.Type = "Purchase"
	s := success.Success{}
	s.Success = true
	successJson, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}
	bitcoinRepository.Create(b)
	jsonResponse.ToJson(w, successJson)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	b := bitcoinModel.Bitcoin{}
	err := decoder.Decode(&b)

	if err != nil {
		panic(err)
	}
	b.Total = float64(b.Quantity) * apiBitcoin.DecodeData()
	b.Type = "Sale"
	s := success.Success{}
	s.Success = true
	successJson, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}
	bitcoinRepository.Create(b)
	jsonResponse.ToJson(w, successJson)
}

func GetByUser(w http.ResponseWriter, r *http.Request) {
	bitcoins := bitcoinRepository.GetByUser()

	bitcoinsJson, err := json.Marshal(bitcoins)
	if err != nil {
		panic(err)
	}
	jsonResponse.ToJson(w, bitcoinsJson)
}

func GetByDay(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	firstday := vars.Get("initialday")
	secondday := vars.Get("secondday")

	bitcoins := bitcoinRepository.GetByDay(firstday, secondday)

	bitcoinsJson, err := json.Marshal(bitcoins)
	if err != nil {
		panic(err)
	}
	jsonResponse.ToJson(w, bitcoinsJson)
}
