package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")

	response := Response{
		Message: fmt.Sprintf("Stay home, stay safe %s", name),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("hey there, log here!")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
