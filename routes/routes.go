package routes

import (
	"net/http"

	"github.com/MockApis/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	// Sila transact endpoint
	r.HandleFunc("/sila_transact", services.HandleSilaTransact).Methods(http.MethodPost)

	return r
}
