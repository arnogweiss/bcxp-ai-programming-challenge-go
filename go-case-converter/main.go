package main

import (
	"fmt"
	"net/http"

	"go-case-converter/handlers"
)

func main() {
	http.HandleFunc("/convert/camel-to-snake", handlers.CamelToSnakeHandler)
	http.HandleFunc("/convert/snake-to-camel", handlers.SnakeToCamelHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
