package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MockApis/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Println("WARNING: WEBHOOK_URL is not set")
	} else {
		log.Println("Webhook URL:", webhookURL)
	}

	r := routes.RegisterRoutes()

	log.Println("🚀 Mock Sila API running on :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
