package model

import (
	"time"

	grpcServerV1 "github.com/aprln/trading-chart/grpc-server/pkg/grpcserver/v1"
)

type OHLCCandlestick struct {
	Symbol     string
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	StartTime  time.Time
	CreatedAt  time.Time
}

func (o OHLCCandlestick) FromStreamOHLCCandlestickRequest(req *grpcServerV1.StreamOHLCCandlestickRequest) OHLCCandlestick {
	return OHLCCandlestick{
		Symbol:     req.Symbol,
		OpenPrice:  req.GetOpenPrice(),
		HighPrice:  req.GetHighPrice(),
		LowPrice:   req.GetLowPrice(),
		ClosePrice: req.GetClosePrice(),
		StartTime:  time.UnixMilli(req.GetStartTime()),
	}
}
