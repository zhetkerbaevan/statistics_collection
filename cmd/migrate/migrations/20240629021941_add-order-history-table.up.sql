CREATE TABLE IF NOT EXISTS order_history (
    id SERIAL PRIMARY KEY,
    client_name VARCHAR(255) NOT NULL,
    exchange_name VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL,
    pair VARCHAR(255) NOT NULL,
    side VARCHAR(255) NOT NULL,
    types VARCHAR(255) NOT NULL,
    base_qty FLOAT8 NOT NULL,
    price FLOAT8 NOT NULL,
    algorithm_name_placed VARCHAR(255),
    lowest_sell_prc FLOAT8,
    highest_buy_prc FLOAT8,
    commission_quote_qty FLOAT8,
    time_placed TIMESTAMPTZ NOT NULL
);