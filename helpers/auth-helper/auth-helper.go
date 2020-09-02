package authhelper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	errorModel "api-golang/models/error"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		w.Header().Set("Content-Type", "application/json")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					w.WriteHeader(http.StatusBadRequest)
					messages := []errorModel.Message{errorModel.Message{Pt: error.Error(), En: error.Error()}}
					json.NewEncoder(w).Encode(errorModel.Error{Code: 1, Messages: messages})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					messages := []errorModel.Message{errorModel.Message{Pt: "Token de autorização inválido", En: "Invalid authorization token"}}
					json.NewEncoder(w).Encode(errorModel.Error{Code: 1, Messages: messages})
				}
			} else if len(bearerToken) == 1 {
				w.WriteHeader(http.StatusBadRequest)
				messages := []errorModel.Message{errorModel.Message{Pt: "O token de autorização debe seguir o seguinte formato: Bearer {{token}}", En: "The authorization token must have the following form: Bearer {{token}}"}}
				json.NewEncoder(w).Encode(errorModel.Error{Code: 1, Messages: messages})
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			messages := []errorModel.Message{errorModel.Message{Pt: "O header de autorização é obrigatório", En: "An authorization header is required"}}
			json.NewEncoder(w).Encode(errorModel.Error{Code: 2, Messages: messages})
		}
	})
}
