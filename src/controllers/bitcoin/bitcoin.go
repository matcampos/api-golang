package bitcoinController

import (
	"encoding/json"
	"net/http"

	jsonResponse "../../helpers/json-response"
	bitcoinModel "../../models/bitcoin"
	bitcoinRepository "../../repository/bitcoin"
)

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	b := bitcoinModel.Bitcoin{}
	err := decoder.Decode(&b)

	if err != nil {
		panic(err)
	}

	bitcoinJson, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}
	bitcoinRepository.Create(b)
	jsonResponse.ToJson(w, bitcoinJson)
}

// func GetAll(w http.ResponseWriter, r *http.Request) {
// 	users := userRepository.GetAll()

// 	usersJson, err := json.Marshal(users)
// 	if err != nil {
// 		panic(err)
// 	}
// 	jsonResponse.ToJson(w, usersJson)
// }
