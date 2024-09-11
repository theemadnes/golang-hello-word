package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Now you can access variables like before
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable not set; setting to default of 8080")
		port = "8080"
	}
	fmt.Println("PORT:", port)

	// Start the HTTP server

	http.HandleFunc("/", handleRoot)

	fmt.Println("Server listening on port ", port)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// See if MSG environment variable is set, and if not set to default
	msg := os.Getenv("MSG")
	if msg == "" {
		msg = "Hello from Go!"
	}

	// Create a data structure to represent the JSON response
	responseData := map[string]string{
		"message": msg,
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Encode the data as JSON and write it to the response
	json.NewEncoder(w).Encode(responseData)
}
