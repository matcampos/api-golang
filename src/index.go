package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    routes "./routes"
)

func main(){
    var router = mux.NewRouter()

	router.HandleFunc("/healthcheck", routes.HealthCheck).Methods("GET")
    router.HandleFunc("/message", routes.HandleQryMessage).Methods("GET")
    router.HandleFunc("/m/{msg}", routes.HandleUrlMessage).Methods("GET")
    router.HandleFunc("/m", routes.HandlePost).Methods("POST")
    router.HandleFunc("/users", routes.HandleGet).Methods("GET")
    fmt.Println("Running server on port 3000")
    log.Fatal(http.ListenAndServe(":3000", router))
    
}
