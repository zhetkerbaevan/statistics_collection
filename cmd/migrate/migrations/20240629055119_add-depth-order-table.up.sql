CREATE TABLE IF NOT EXISTS depth_order (
    id SERIAL PRIMARY KEY,
    orderbook_id INT REFERENCES order_book(id),
    price FLOAT8 NOT NULL,
    base_qty FLOAT8 NOT NULL,
    order_type VARCHAR(4) CHECK (order_type IN ('ask', 'bid')) NOT NULL
);