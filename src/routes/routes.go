package routes

import (
	bitcoinController "../controllers/bitcoin"
	userController "../controllers/user"
	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	// authenticated := mux.NewRouter()
	// mw := jwtmiddleware.New(jwtmiddleware.Options{
	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("secret"), nil
	// 	},
	// 	SigningMethod: jwt.SigningMethodHS256,
	// })

	// an := negroni.New(
	// 	negroni.HandlerFunc(mw.HandlerWithNext),
	// 	negroni.Wrap(authenticated))

	// n := negroni.Classic()
	// router.PathPrefix("/api").Handler(an)
	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/api/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/api/bitcoins", bitcoinController.CreatePurchase).Methods("POST")
	router.HandleFunc("/api/bitcoins/sale", bitcoinController.CreateSale).Methods("POST")
	router.HandleFunc("/api/bitcoins/getbyuser", bitcoinController.GetByUser).Methods("GET")
	router.HandleFunc("/api/bitcoins/getbyday/{day}", bitcoinController.GetByDay).Methods("GET")
	// router.HandleFunc("/healthcheck", routes.HealthCheck).Methods("GET")
	// router.HandleFunc("/message", routes.HandleQryMessage).Methods("GET")
	// router.HandleFunc("/m/{msg}", routes.HandleUrlMessage).Methods("GET")
	// router.HandleFunc("/m", routes.HandlePost).Methods("POST")

	// n.UseHandler(authenticated)
	return router
}
