package orderbook

import (
	"database/sql"

	"github.com/zhetkerbaevan/statistics_collection/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetOrderBook(exchange_name, pair string) ([]*models.DepthOrder, error) {
	return nil, nil
}

func (s *Store) SaveOrderBook(exchange_name, pair string, orderBook []*models.DepthOrder) error {
	return nil
}

func scanRowsIntoOrderBook(rows *sql.Rows) {
	
}

