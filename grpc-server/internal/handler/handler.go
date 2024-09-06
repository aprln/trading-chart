package handler

import (
	"fmt"

	"github.com/aprln/trading-chart/grpc-server/internal/model"
	v1 "github.com/aprln/trading-chart/grpc-server/pkg/grpcserver/v1"
	"google.golang.org/grpc"
)

type service interface {
	CreateOHLCCandlestick(ohlc model.OHLCCandlestick) error
}

func New(svc service) *Handler {
	return &Handler{
		service: svc,
	}
}

// Handler represents struct that implements TradingChartServer.
type Handler struct {
	service service
	v1.UnsafeTradingChartServer
}

func (h Handler) StreamOHLCCandlestick(
	req *v1.StreamOHLCCandlestickRequest,
	stream grpc.ServerStreamingServer[v1.StreamOHLCCandlestickResponse],
) error {
	if err := h.service.CreateOHLCCandlestick(model.OHLCCandlestick{}.FromStreamOHLCCandlestickRequest(req)); err != nil {
		return fmt.Errorf("create OHLC Candlestick: %w", err)
	}

	if err := stream.Send(&v1.StreamOHLCCandlestickResponse{}); err != nil {
		return fmt.Errorf("send OHLC Candlestick response: %w", err)
	}

	return nil
}
