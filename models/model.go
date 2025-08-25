package models

// SilaTransactRequest represents the incoming transact request payload
type SilaTransactRequest struct {
	Amount        int64  `json:"amount"`
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
}

// SilaTransactResponse represents the response after transact
type SilaTransactResponse struct {
	// Reference       string `json:"reference"`
	SilaReferenceID string `json:"sila_reference_id"`
	Message         string `json:"message"`
	Success         bool   `json:"success"`
	Status          string `json:"status"`
	ResponseTimeMS  string `json:"response_time_ms"`
	TransactionID   string `json:"transaction_id"`
	Itinerary       string `json:"itinerary_selected"`
	Description     string `json:"description,omitempty"`
}

// WebhookEvent represents the webhook payload
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
		// ReturnCode      string `json:"return_code"`
		// ReturnDesc      string `json:"return_description"`
		// OldGraph        string `json:"old_graph"`
		// OldRoute        string `json:"old_route"`
		// NewRoute        string `json:"new_route"`
	} `json:"event_details"`
}
