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
		panic(err)
	}

	userJson, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}
	userRepository.Create(u)
	jsonResponse.ToJson(w, userJson)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	users := userRepository.GetAll()

	usersJson, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	jsonResponse.ToJson(w, usersJson)
}
