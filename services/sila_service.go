package services

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/MockApis/models"

	"github.com/google/uuid"
)

func HandleSilaTransact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SilaTransactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		log.Println("Decode error:", err)
		return
	}

	transactionID := uuid.New().String()
	response := models.SilaTransactResponse{
		SilaReferenceID: uuid.New().String(),
		Message:         "Transaction submitted to the processing queue.",
		Success:         true,
		Status:          "SUCCESS",
		ResponseTimeMS:  "171",
		TransactionID:   transactionID,
		Itinerary:       "STANDARD_ACH",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Launch webhook asynchronously
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Println("WEBHOOK_URL not set, skipping webhook")
		return
	}

	go triggerWebhook(webhookURL, transactionID, req.Amount)
}

func triggerWebhook(webhookURL, transactionID string, amount string) {
	time.Sleep(15 * time.Second) // delay ~1 min

	event := models.WebhookEvent{
		EventTime: time.Now().Format(time.RFC3339),
		EventType: "transaction",
		EventUUID: uuid.New().String(),
	}
	event.EventDetail.Transaction = transactionID
	event.EventDetail.TransactionType = "transfer"
	event.EventDetail.SilaAmount = amount

	outcomes := []string{"success", "failed", "review", "refunded", "refund_failed"}
	event.EventDetail.Outcome = outcomes[rand.Intn(len(outcomes))]

	// event.EventDetail.Entity = userHandle
	event.EventDetail.ProcessingType = "STANDARD_ACH"

	body, _ := json.Marshal(event)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending webhook to %s: %v", webhookURL, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Webhook event sent to %s with status: %s", webhookURL, resp.Status)
}
