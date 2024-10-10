package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, World!"); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
