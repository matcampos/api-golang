package jsonResponse

import (
	"encoding/json"
	"net/http"

	errorModel "../../models/error"
	"github.com/go-errors/errors"
	"github.com/go-sql-driver/mysql"
)

func ToJson(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func ToError(w http.ResponseWriter, errorStack error, status int) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	/*
		Errors treatment
	*/
	if err, ok := errorStack.(*mysql.MySQLError); ok {
		messages := []errorModel.Message{errorModel.Message{Pt: err.Message, En: err.Message}}
		errors := errorModel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
		return
	}

	messages := []errorModel.Message{errorModel.Message{Pt: errorStack.(*errors.Error).ErrorStack(), En: errorStack.(*errors.Error).ErrorStack()}}
	errors := errorModel.Error{Code: http.StatusInternalServerError, Messages: messages}

	json.NewEncoder(w).Encode(errors)
}

func CustomError(w http.ResponseWriter, errorM errorModel.Error, status int) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(errorM)
}
