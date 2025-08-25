package services

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/MockApis/models"

	"github.com/google/uuid"
)

// HandleSilaTransact handles the /sila_transact route
func HandleSilaTransact(w http.ResponseWriter, r *http.Request) {
	var req models.SilaTransactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	transactionID := uuid.New().String()
	response := models.SilaTransactResponse{
		Reference:       req.Header.Reference,
		SilaReferenceID: uuid.New().String(),
		Message:         "Transaction submitted to the processing queue.",
		Success:         true,
		Status:          "SUCCESS",
		ResponseTimeMS:  "171",
		TransactionID:   transactionID,
		Itinerary:       "STANDARD_ACH",
		Description:     req.Description,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Launch webhook trigger asynchronously
	go triggerWebhook(req.Header.UserHandle, transactionID, req.Amount)
}

// triggerWebhook simulates a webhook after ~1 minute
func triggerWebhook(userHandle, transactionID string, amount int64) {
	time.Sleep(60 * time.Second) // wait 1 minute

	event := models.WebhookEvent{
		EventTime: time.Now().Unix(),
		EventType: "transaction",
		EventUUID: uuid.New().String(),
	}
	event.EventDetail.Transaction = transactionID
	event.EventDetail.TransactionType = "transfer"
	event.EventDetail.SilaAmount = amount

	// Random outcome (for demo) → could be success/failed/review
	outcomes := []string{"success", "failed", "review", "refunded", "refund_failed"}
	event.EventDetail.Outcome = outcomes[rand.Intn(len(outcomes))]

	event.EventDetail.Entity = userHandle
	event.EventDetail.ProcessingType = "STANDARD_ACH"

	body, _ := json.Marshal(event)

	resp, err := http.Post("http://localhost:5000/webhook_event_receiver",
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error sending webhook: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Webhook event sent with status:", resp.Status)
}
