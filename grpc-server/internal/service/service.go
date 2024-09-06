package service

import (
	"fmt"

	"github.com/aprln/trading-chart/grpc-server/internal/model"
)

type repository interface {
	CreateOHLCCandlestick(ohlc model.OHLCCandlestick) error
}

type service struct {
	repo repository
}

func New(repo repository) *service {
	return &service{
		repo: repo,
	}
}

func (s service) CreateOHLCCandlestick(ohlc model.OHLCCandlestick) error {
	if err := s.repo.CreateOHLCCandlestick(ohlc); err != nil {
		return fmt.Errorf("cannot create OHLC Candlestick: %w", err)
	}

	return nil
}
