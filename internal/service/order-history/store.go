package orderhistory

import (
	"database/sql"

	"github.com/zhetkerbaevan/statistics_collection/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db : db}
}

func (s *Store) GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	rows, err := s.db.Query("SELECT * FROM order_history WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4", client.Client_name, client.Exchange_name, client.Label, client.Pair)
	if err != nil {
		return nil, err
	}

	histories := make([]*models.HistoryOrder, 0)
	for rows.Next() {
		h, err := scanRowsIntoOrderHistory(rows)
		if err != nil {
			return nil, err
		}
		histories = append(histories, h)
	}
	
	return histories, nil
}

func (s *Store) SaveOrder(client *models.Client, order *models.HistoryOrder) error {
	_, err := s.db.Exec("INSERT INTO order_history (client_name, exchange_name, label, pair, side, types, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", client.Client_name, client.Exchange_name, client.Label, client.Pair, order.Side, order.Types, order.Base_qty, order.Price, order.Algorithm_name_placed, order.Lowest_sell_prc, order.Highest_buy_prc, order.Commission_quote_qty, order.Time_placed)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoOrderHistory(rows *sql.Rows) (*models.HistoryOrder, error) {
	history := &models.HistoryOrder{}
	err := rows.Scan(
		&history.Id,
		&history.Client_name,
		&history.Exchange_name,
		&history.Label,
		&history.Pair,
		&history.Side,
		&history.Types,
		&history.Base_qty,
		&history.Price,
		&history.Algorithm_name_placed,
		&history.Lowest_sell_prc,
		&history.Highest_buy_prc,
		&history.Commission_quote_qty,
		&history.Time_placed,
	)
	if err != nil {
		return nil, err
	}
	return history, nil
}
