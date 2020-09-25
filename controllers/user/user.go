package userController

import (
	"encoding/json"
	"net/http"

	jsonResponse "api-golang/helpers/json-response"
	errorModel "api-golang/models/error"
	userModel "api-golang/models/user"
	userRepository "api-golang/repository/user"

	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	u := userModel.User{}
	err := decoder.Decode(&u)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	invalidUser := u.Validate()

	if invalidUser != nil {
		errorMessage := invalidUser.Error()
		messages := []errorModel.Message{errorModel.Message{Pt: errorMessage, En: errorMessage}}
		err := errorModel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonResponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)

	if err != nil {
		jsonResponse.ToError(w, err, 0)
		return
	}

	u.Password = string(hash)

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
