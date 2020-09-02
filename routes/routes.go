package routes

import (
	bitcoinController "api-golang/controllers/bitcoin"
	userController "api-golang/controllers/user"
	authHelper "api-golang/helpers/auth-helper"
	generateToken "api-golang/helpers/generate-token"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", userController.Create).Methods("POST")
	router.HandleFunc("/api/bitcoins", authHelper.Authenticate(bitcoinController.CreatePurchase)).Methods("POST")
	router.HandleFunc("/api/bitcoins/sale", authHelper.Authenticate(bitcoinController.CreateSale)).Methods("POST")
	router.HandleFunc("/api/bitcoins/getbyuser", authHelper.Authenticate(bitcoinController.GetByUser)).Methods("GET")
	router.HandleFunc("/api/bitcoins/getbyday", authHelper.Authenticate(bitcoinController.GetByDay)).Methods("GET")
	router.HandleFunc("/authenticate", generateToken.CreateToken).Methods("POST")
	return router
}
