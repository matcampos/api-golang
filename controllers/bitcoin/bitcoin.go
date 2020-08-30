package bitcoinController

import (
	"encoding/json"
	"net/http"

	jsonResponse "../../helpers/json-response"
	"../../models/apiBitcoin"
	bitcoinModel "../../models/bitcoin"
	errorModel "../../models/error"
	success "../../models/success"
	bitcoinRepository "../../repository/bitcoin"
)

func CreatePurchase(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	b := bitcoinModel.Bitcoin{}
	err := decoder.Decode(&b)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	price, err := apiBitcoin.DecodeData()

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	b.Total = float64(b.Quantity) * price
	b.Type = "Purchase"
	s := success.Success{Success: true}
	successJson, err := json.Marshal(s)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	invalidBitcoin := b.ValidateBitcoinStruct()

	if invalidBitcoin != nil {
		errorMessage := invalidBitcoin.Error()
		messages := []errorModel.Message{errorModel.Message{Pt: errorMessage, En: errorMessage}}
		err := errorModel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonResponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	errorr := bitcoinRepository.Create(b)

	if errorr != nil {
		jsonResponse.ToError(w, errorr, 0)
		return
	}

	jsonResponse.ToJson(w, successJson)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	b := bitcoinModel.Bitcoin{}
	err := decoder.Decode(&b)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	price, err := apiBitcoin.DecodeData()

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	b.Total = float64(b.Quantity) * price
	b.Type = "Sale"
	s := success.Success{}
	s.Success = true
	successJson, err := json.Marshal(s)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	invalidBitcoin := b.ValidateBitcoinStruct()

	if invalidBitcoin != nil {
		errorMessage := invalidBitcoin.Error()
		messages := []errorModel.Message{errorModel.Message{Pt: errorMessage, En: errorMessage}}
		err := errorModel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonResponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	errorr := bitcoinRepository.Create(b)

	if errorr != nil {
		jsonResponse.ToError(w, errorr, 0)
		return
	}

	jsonResponse.ToJson(w, successJson)
}

func GetByUser(w http.ResponseWriter, r *http.Request) {
	bitcoins, err := bitcoinRepository.GetByUser()

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	bitcoinsJson, err := json.Marshal(bitcoins)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	jsonResponse.ToJson(w, bitcoinsJson)
}

func GetByDay(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	firstday := vars.Get("initialday")
	secondday := vars.Get("secondday")

	bitcoins, err := bitcoinRepository.GetByDay(firstday, secondday)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	bitcoinsJSON, err := json.Marshal(bitcoins)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	jsonResponse.ToJson(w, bitcoinsJSON)
}
