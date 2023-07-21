package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PingResponse struct {
	Ping     string `json:"ping"`
	Response string `json:"response"`
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{
		Ping:     "ping",
		Response: "pong",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", pingHandler)
	fmt.Println("Server is listening on http://localhost:8080")
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
