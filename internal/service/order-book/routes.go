package orderbook

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/zhetkerbaevan/statistics_collection/internal/models"
	"github.com/zhetkerbaevan/statistics_collection/internal/utils"
)

type Handler struct { // Handler shouldn't know exactly how user operations are implemented
	//It just uses this interface and do ops
	store models.OrderBookStoreInterface 
}

func NewHandler(store models.OrderBookStoreInterface) *Handler { //Creates instance of Handler
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orderbook/{exchange_name}/{pair}", h.handleGetOrderBook).Methods("GET")
	router.HandleFunc("/orderbook", h.handleSaveOrderBook).Methods("POST")
}

func (h *Handler) handleGetOrderBook(w http.ResponseWriter, r *http.Request) {
	//Extract URL parameters
	vars := mux.Vars(r)
	exchangeName := vars["exchange_name"]
	pair := vars["pair"]

	depth_orders, err := h.store.GetOrderBook(exchangeName, pair)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, depth_orders)
}

func (h *Handler) handleSaveOrderBook(w http.ResponseWriter, r *http.Request) {
	var payload models.OrderBook
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("INVALID PAYLOAD %v", errors))
		return
	}
	err := h.store.SaveOrderBook(payload.Exchange_name, payload.Pair, payload.Depth_orders)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}
