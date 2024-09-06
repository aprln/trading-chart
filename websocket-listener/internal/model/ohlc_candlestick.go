package model

import (
	"time"
)

type OHLCCandlestick struct {
	Symbol     string
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	StartTime  time.Time
}
