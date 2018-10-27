package routes

import (
	bitcoinController "../controllers/bitcoin"
	userController "../controllers/user"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/bitcoins", bitcoinController.Create).Methods("POST")
	// router.HandleFunc("/healthcheck", routes.HealthCheck).Methods("GET")
	// router.HandleFunc("/message", routes.HandleQryMessage).Methods("GET")
	// router.HandleFunc("/m/{msg}", routes.HandleUrlMessage).Methods("GET")
	// router.HandleFunc("/m", routes.HandlePost).Methods("POST")
	return router
}
