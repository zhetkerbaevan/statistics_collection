package orderhistory

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
	store models.OrderHistoryInterface
}

func NewHandler(store models.OrderHistoryInterface) *Handler { //Creates instance of Handler
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orderhistory/{client_name}/{exchange_name}/{label}/{pair}", h.handleGetOrderHistory).Methods("GET")
	router.HandleFunc("/orderhistory", h.handleSaveOrderHistory).Methods("POST")
}

func (h *Handler) handleGetOrderHistory(w http.ResponseWriter, r *http.Request) {
	//Extract URL parameters
	vars := mux.Vars(r)
	clientName := vars["client_name"]
	exchangeName := vars["exchange_name"]
	label := vars["label"]
	pair := vars["pair"]

	//Create a Client structure for passing to store.GetOrderHistory
	client := &models.Client{
		Client_name:   clientName,
		Exchange_name: exchangeName,
		Label:         label,
		Pair:          pair,
	}

	//Retrieve data from the database
	histories, err := h.store.GetOrderHistory(client)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Return data in JSON format
	utils.WriteJSON(w, http.StatusOK, histories)
}


func (h *Handler) handleSaveOrderHistory(w http.ResponseWriter, r *http.Request) {
	var payload models.HistoryOrder //Payload is the data that we receive/send
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("INVALID PAYLOAD %v", errors))
		return

	}

	client := &models.Client{
		Client_name:   payload.Client_name, // Assuming payload.ClientName matches client.ClientName
		Exchange_name: payload.Exchange_name,
		Label:        payload.Label,
		Pair:         payload.Pair,
	}
	
	err := h.store.SaveOrder(client, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}