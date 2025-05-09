package main

import (
	"fmt"
	"net/http"

	"go-case-converter/handlers"
)

func main() {
	http.HandleFunc("/convert/camel-to-snake", handlers.CamelToSnakeHandler)
	http.HandleFunc("/convert/snake-to-camel", handlers.SnakeToCamelHandler)
	http.HandleFunc("/convert/kebab-to-camel", handlers.KebabToCamelHandler)
	http.HandleFunc("/convert/kebab-to-snake", handlers.KebabToSnakeHandler)
	http.HandleFunc("/convert/camel-to-kebab", handlers.CamelToKebabHandler)
	http.HandleFunc("/convert/snake-to-kebab", handlers.SnakeToKebabHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
