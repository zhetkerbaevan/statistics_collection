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
	var id int
    err := s.db.QueryRow("SELECT id FROM order_book WHERE exchange = $1 AND pair = $2", exchange_name, pair).Scan(id) 
    if err != nil {
        return nil, err
    }

	depth_orders, err := s.GetDepthOrderById(id)
	if err != nil {
        return nil, err
    }

	return depth_orders, nil
}



func (s *Store) SaveOrderBook(exchange_name, pair string, orderBook []*models.DepthOrder) error {
	var orderBookID int
	//Insert data to table order_book returning id of this row
	err := s.db.QueryRow(
		`INSERT INTO order_book (exchange, pair) VALUES ($1, $2) RETURNING id`,
		exchange_name, pair,
	).Scan(&orderBookID)
	if err != nil {
		return err
	}

	//Insert data to depth_order
	for _, order := range orderBook {
		_, err := s.db.Exec(
			`INSERT INTO depth_order (orderbook_id, price, base_qty, order_type) VALUES ($1, $2, $3, $4)`,
			orderBookID, order.Price, order.BaseQty, order.OrderType,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) GetDepthOrderById(id int) ([]*models.DepthOrder, error) {
    rows, err := s.db.Query("SELECT * FROM depth_order WHERE orderbook_id = $1", id)
    if err != nil {
        return nil, err
    }

    depth_orders := make([]*models.DepthOrder, 0)
    for rows.Next() {
        d, err := scanRowsIntoDepthOrder(rows)
        if err != nil {
            return nil, err
        }
        depth_orders = append(depth_orders, &models.DepthOrder{Price: d.Price, BaseQty: d.BaseQty})
    }
    return depth_orders, nil
}


func scanRowsIntoDepthOrder(rows *sql.Rows) (*models.DepthOrderData, error) {
	d := new(models.DepthOrderData)
	err := rows.Scan(
		&d.Id,
		&d.OrderBookId,
		&d.Price,
		&d.BaseQty,
		&d.OrderType,
	)
	if err != nil {
		return nil, err
	}
	return d, nil
}
