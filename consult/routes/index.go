package routes

import (
	routes "./routes"
	"github.com/gorilla/mux"
)

func Index() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", routes.HealthCheck).Methods("GET")
	router.HandleFunc("/message", routes.HandleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", routes.HandleUrlMessage).Methods("GET")
	router.HandleFunc("/m", routes.HandlePost).Methods("POST")
	router.HandleFunc("/users", routes.HandleGet).Methods("GET")
	return router
}
