package main

import (
	"log"
	"net/http"

	"github.com/MockApis/routes"
)

func main() {
	r := routes.RegisterRoutes()

	log.Println("🚀 Mock Sila API running on :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
