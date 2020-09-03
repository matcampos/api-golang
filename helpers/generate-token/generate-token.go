package generateToken

import (
	con "api-golang/config/db"
	jsonResponse "api-golang/helpers/json-response"
	ErrorModel "api-golang/models/error"
	loginModel "api-golang/models/login"
	UserModel "api-golang/models/user"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(w http.ResponseWriter, req *http.Request) {
	login := loginModel.Login{}
	_ = json.NewDecoder(req.Body).Decode(&login)

	invalidBody := login.Validate()
	if invalidBody != nil {
		errorMessage := invalidBody.Error()
		messages := []ErrorModel.Message{ErrorModel.Message{Pt: errorMessage, En: errorMessage}}
		err := ErrorModel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonResponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	db, err := con.Connect()
	if err != nil {
		return
	}

	userModel := UserModel.User{}

	row := db.QueryRow("select * from user where email=?", login.Email)

	db.Close()

	switch err := row.Scan(&userModel.Id, &userModel.Name, &userModel.Email, &userModel.Password, &userModel.Birthdate); err {
	case sql.ErrNoRows:
		messages := []ErrorModel.Message{ErrorModel.Message{
			Pt: "Usuário não encontrado",
			En: "User not found."}}
		customError := ErrorModel.Error{
			Code:     http.StatusNotFound,
			Messages: messages,
			Field:    "email"}
		jsonResponse.CustomError(w, customError, http.StatusNotFound)
		return
	case nil:

		byteArrayPass := []byte(login.Password)
		findedUserByteArray := []byte(userModel.Password)

		invalidPassword := bcrypt.CompareHashAndPassword(findedUserByteArray, byteArrayPass)

		if invalidPassword != nil {
			messages := []ErrorModel.Message{ErrorModel.Message{
				Pt: "Senha incorreta.",
				En: "Invalid password"}}
			customError := ErrorModel.Error{
				Code:     http.StatusForbidden,
				Messages: messages,
				Field:    "password"}
			jsonResponse.CustomError(w, customError, http.StatusForbidden)
			return
		}

		clains := loginModel.CustomClain{
			userModel.Name,
			userModel.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				Issuer:    "test",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)

		singingKey := []byte(os.Getenv("API_SECRET"))

		tokenString, err := token.SignedString(singingKey)

		if err != nil {
			jsonResponse.ToError(w, err, 0)
			return
		}

		response, err := json.Marshal(loginModel.JwtToken{Token: tokenString})

		if err != nil {
			jsonResponse.ToError(w, err, 0)
			return
		}

		jsonResponse.ToJson(w, response)
	default:
		panic(err)
	}
}
