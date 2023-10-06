package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX level error:", msg)
	}
	type errorResponse struct {
		Error string `json:"error"` //this is to convert json object
	}
	respondWIthJson(w, code, errorResponse{
		Error: msg,
	})
}

func respondWIthJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal resposnse: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("COntent-type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
