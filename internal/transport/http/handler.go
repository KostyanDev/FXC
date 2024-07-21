package http

import (
	"net/http"
)

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	if h.service.SubscribeService(r.Context(), address) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Subscribed to " + address))
	} else {
		http.Error(w, "already subscribed", http.StatusBadRequest)
	}
}
