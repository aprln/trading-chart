package model

import (
	"strconv"
	"time"

	binanceConnector "github.com/binance/binance-connector-go"
)

type AggTrade struct {
	Symbol    string
	Price     float64
	TradeTime time.Time
}

func (at AggTrade) FromWsAggTradeEvent(event *binanceConnector.WsAggTradeEvent) (AggTrade, error) {
	price, err := at.parsePrice(event.Price)
	if err != nil {
		return AggTrade{}, err
	}

	return AggTrade{
		TradeTime: time.UnixMilli(event.TradeTime),
		Symbol:    event.Symbol,
		Price:     price,
	}, nil

}

func (at AggTrade) parsePrice(priceStr string) (float64, error) {
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
