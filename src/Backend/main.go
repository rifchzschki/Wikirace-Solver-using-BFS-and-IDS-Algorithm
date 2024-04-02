package main

import (
	"encoding/json"
	"net/http"
)

func main() {
    http.HandleFunc("/api/data", dataHandler)
    http.ListenAndServe(":8080", nil)
}

type ApiResponse struct {
    Message string `json:"message"`
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    // Process the request and send back a response
    response := ApiResponse{Message: "Hello from the Golang API!"}
    json.NewEncoder(w).Encode(response)
}