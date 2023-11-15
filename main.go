package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello, World!",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
