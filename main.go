package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "asciiart/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Request)
	fmt.Println("Server is starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
}
