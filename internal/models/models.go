package models

import "time"

type DepthOrder struct {
	Price   float64 `json:"price"`
	BaseQty float64 `json:"baseqty"`
}

type HistoryOrder struct {
	Id int `json:"id"`
	Client_name           string `json:"client_name"`
	Exchange_name         string `json:"exchange_name"`
	Label                 string `json:"label"`
	Pair                  string `json:"pair"`
	Side                  string `json:"side"`
	Types                 string `json:"types"`
	Base_qty              float64 `json:"base_qty"`
	Price                 float64 `json:"price"`
	Algorithm_name_placed string `json:"algorithm_name_placed"`
	Lowest_sell_prc       float64 `json:"lowest_sell_prc"`
	Highest_buy_prc       float64 `json:"highest_buy_prc"`
	Commission_quote_qty  float64 `json:"commission_quote_qty"`
	Time_placed           time.Time `json:"time_placed"`
}

type Client struct {
	Client_name   string `json:"client_name"`
	Exchange_name string `json:"exchange_name"`
	Label         string `json:"label"`
	Pair          string `json:"pair"`
}

type OrderBookStoreInterface interface {
	GetOrderBook(exchange_name, pair string) ([]*DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*DepthOrder) error
}

type OrderHistoryInterface interface {
	GetOrderHistory(client *Client)  ([]*HistoryOrder, error)
	SaveOrder(client *Client, order *HistoryOrder) error
}