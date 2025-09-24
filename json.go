package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal JSON Response: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
