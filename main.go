package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MockApis/services"
)

func main() {
	http.HandleFunc("/sila_transact", services.HandleSilaTransact)
	http.HandleFunc("/webhook_event_receiver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Webhook received")
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Mock Sila API running on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
