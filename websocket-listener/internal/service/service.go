package service

import (
	"context"
	"log"
	"sync"
	"time"

	v1 "github.com/aprln/trading-chart/grpc-server/pkg/grpcserver/v1"
	"github.com/aprln/trading-chart/websocket-listener/internal/model"
)

func New(tradingChartClient v1.TradingChartClient) *service {
	return &service{
		tradingChartClient: tradingChartClient,
	}
}

type service struct {
	tradingChartClient v1.TradingChartClient

	curOHLC *model.OHLCCandlestick

	sm sync.Map
}

func (s *service) HandleAggTrade(ctx context.Context, aggTrade model.AggTrade) error {
	ohlc, shouldSend := s.aggregateOHLCCandlesticks(aggTrade)

	if shouldSend {
		if _, err := s.tradingChartClient.StreamOHLCCandlestick(
			ctx, &v1.StreamOHLCCandlestickRequest{
				Symbol:     ohlc.Symbol,
				OpenPrice:  ohlc.OpenPrice,
				HighPrice:  ohlc.HighPrice,
				LowPrice:   ohlc.LowPrice,
				ClosePrice: ohlc.ClosePrice,
				StartTime:  ohlc.StartTime.UnixMilli(),
			},
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) aggregateOHLCCandlesticks(aggTrade model.AggTrade) (result model.OHLCCandlestick, shouldSend bool) {
	if s.curOHLC == nil {
		// First candle
		s.curOHLC = toPointer(s.newOHLCCandlestick(aggTrade))

		return *s.curOHLC, false
	}

	if s.curOHLC.Symbol != aggTrade.Symbol {
		log.Fatal("bad")
	}

	if !aggTrade.TradeTime.Truncate(time.Minute).Equal(s.curOHLC.StartTime) {
		// New candle.
		result = *s.curOHLC

		s.curOHLC = toPointer(s.newOHLCCandlestick(aggTrade))

		return result, true
	}

	if aggTrade.Price < s.curOHLC.LowPrice {
		s.curOHLC.LowPrice = aggTrade.Price
	}

	if aggTrade.Price > s.curOHLC.HighPrice {
		s.curOHLC.HighPrice = aggTrade.Price
	}

	s.curOHLC.ClosePrice = aggTrade.Price

	return *s.curOHLC, false
}

func (s *service) newOHLCCandlestick(aggTrade model.AggTrade) model.OHLCCandlestick {
	return model.OHLCCandlestick{
		Symbol:     aggTrade.Symbol,
		OpenPrice:  aggTrade.Price,
		HighPrice:  aggTrade.Price,
		LowPrice:   aggTrade.Price,
		ClosePrice: aggTrade.Price,
		StartTime:  aggTrade.TradeTime.Truncate(time.Minute),
	}
}

func toPointer[T any](v T) *T {
	return &v
}
