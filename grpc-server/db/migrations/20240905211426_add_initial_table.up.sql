CREATE TABLE IF NOT EXISTS ohlc_candlestick
(
    id          bigint         NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    symbol      varchar(50)    NOT NULL,
    open_price  numeric(18, 8) NOT NULL,
    high_price  numeric(18, 8) NOT NULL,
    low_price   numeric(18, 8) NOT NULL,
    close_price numeric(18, 8) NOT NULL,
    start_time  bigint         NOT NULL,
    created_at  timestamptz    NOT NULL
)