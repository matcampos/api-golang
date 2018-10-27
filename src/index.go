package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"
)

func main() {
	router := routes.Routes()
	fmt.Println("Running server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
