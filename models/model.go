package models

type SilaTransactRequest struct {
	Amount        int64  `json:"amount"`
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
}

type SilaTransactResponse struct {
	SilaReferenceID string `json:"sila_reference_id"`
	Message         string `json:"message"`
	Success         bool   `json:"success"`
	Status          string `json:"status"`
	ResponseTimeMS  string `json:"response_time_ms"`
	TransactionID   string `json:"transaction_id"`
	Itinerary       string `json:"itinerary_selected"`
	Description     string `json:"description,omitempty"`
}

type WebhookEvent struct {
	EventTime   int64  `json:"event_time"`
	EventType   string `json:"event_type"`
	EventUUID   string `json:"event_uuid"`
	EventDetail struct {
		Transaction     string `json:"transaction"`
		TransactionType string `json:"transaction_type"`
		SilaAmount      int64  `json:"sila_amount"`
		Outcome         string `json:"outcome"`
		Entity          string `json:"entity"`
		ProcessingType  string `json:"processing_type"`
		ProviderStatus  string `json:"provider_status"`
	} `json:"event_details"`
}
