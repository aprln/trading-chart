package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/aprln/trading-chart/websocket-listener/internal/model"
	binanceConnector "github.com/binance/binance-connector-go"
)

type service interface {
	HandleAggTrade(ctx context.Context, aggTrade model.AggTrade) error
}

func New(svc service) *Handler {
	return &Handler{
		service: svc,
	}
}

// Handler represents struct that implements TradingChartServer.
type Handler struct {
	service service
}

func (h *Handler) HandleAggTrade(ctx context.Context, event *binanceConnector.WsAggTradeEvent) {
	fmt.Println("EVENT: ", event)

	aggTrade, err := model.AggTrade{}.FromWsAggTradeEvent(event)
	if err != nil {
		log.Printf("Failed to create agg trade model: %v\n", err)

		return
	}

	if err = h.service.HandleAggTrade(ctx, aggTrade); err != nil {
		log.Printf("Failed to handle agg trade: %v\n", err)
	}
}
