package orderbook

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhetkerbaevan/statistics_collection/internal/models"
)

type Handler struct { // Handler shouldn't know exactly how user operations are implemented
	//It just uses this interface and do ops
	store models.OrderBookStoreInterface 
}

func NewHandler(store models.OrderBookStoreInterface) *Handler { //Creates instance of Handler
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orderbook", h.handleGetOrderBook).Methods("GET")
	router.HandleFunc("/orderbook", h.handleSaveOrderBook).Methods("POST")
}

func (h *Handler) handleGetOrderBook(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleSaveOrderBook(w http.ResponseWriter, r *http.Request) {

}