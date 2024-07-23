package http

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/price", handler.GetPricingList).Methods("POST")
}
