package userController

import (
	"encoding/json"
	"net/http"

	jsonResponse "../../helpers/json-response"
	userModel "../../models/user"
	userRepository "../../repository/user"
)

type Bitcoin struct {
	Data string `json:"name"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	u := userModel.User{}
	err := decoder.Decode(&u)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	createdUser, err := userRepository.Create(u)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	userJson, err := json.Marshal(createdUser)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	jsonResponse.ToJson(w, userJson)
}
