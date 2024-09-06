package repository

import (
	"database/sql"
	"time"

	"github.com/aprln/trading-chart/grpc-server/internal/model"
)

func New(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r repository) CreateOHLCCandlestick(ohlc model.OHLCCandlestick) error {
	if _, err := r.db.Exec(
		`INSERT INTO ohlc_candlestick (
            symbol, open_price, high_price, low_price, close_price, start_time, created_at
        )
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		ohlc.Symbol,
		ohlc.OpenPrice,
		ohlc.HighPrice,
		ohlc.LowPrice,
		ohlc.ClosePrice,
		ohlc.StartTime.UnixMilli(),
		time.Now().In(time.UTC),
	); err != nil {
		return err
	}

	return nil
}
