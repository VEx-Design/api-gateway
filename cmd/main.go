package main

import (
	"fmt"
	"log"

	"github.com/Atipat-CMU/api-gateway/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on environment variables")
	}

	// Initialize the router
	r := routes.SetupRouter()

	// Start the server using HTTP
	log.Println("Starting HTTP server on http://localhost:8080")
	err = r.Run(":8080") // Changed from RunTLS to Run
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
