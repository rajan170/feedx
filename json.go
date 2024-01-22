package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, StatusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSON response: %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(StatusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, StatusCode int, msg string) {
	if StatusCode > 499 {
		log.Println("Responded with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, StatusCode, errResponse{
		Error: msg,
	})
}
