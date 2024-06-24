package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "asciiart/handlers"
)


func main() {

	if len(os.Args) !=1 {
		fmt.Println("<Usage>", "<go run .>, <go run main.go>")
		os.Exit(1)
	}
	http.HandleFunc("/", handlers.Request)
	http.HandleFunc("/submit", handlers.Post)
	fmt.Println("Server is starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
}
