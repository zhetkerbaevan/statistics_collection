package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	orderbook "github.com/zhetkerbaevan/statistics_collection/internal/service/order-book"
	orderhistory "github.com/zhetkerbaevan/statistics_collection/internal/service/order-history"
)

type APIServer struct {
	db *sql.DB
	address string
}

func NewAPIServer(db *sql.DB, address string) *APIServer { 
	//Create new instance of APIServer struct
	return &APIServer{
		db : db,
		address: address,
	}
}

func (s *APIServer) Run() error {
	//Create new router
	router := mux.NewRouter()
	
	//Create subrouter (version of api)
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new orderbook store object, passing it the database object
	bookStore := orderbook.NewStore(s.db)
	// Create a new order book service object, passing it the bookStore object
	bookService := orderbook.NewHandler(bookStore)
	
	// Register API routes for working with the order book on the subrouter
	bookService.RegisterRoutes(subrouter)

	historyStore := orderhistory.NewStore(s.db)
	historyService := orderhistory.NewHandler(historyStore)
	historyService.RegisterRoutes(subrouter)
	
	log.Println("Listening on", s.address)

	return http.ListenAndServe(s.address, router)
}