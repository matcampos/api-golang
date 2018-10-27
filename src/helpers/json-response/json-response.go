package jsonResponse

import (
	"net/http"
)

func ToJson(w http.ResponseWriter, personJson []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(personJson)
}
