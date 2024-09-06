package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	v1 "github.com/aprln/trading-chart/grpc-server/pkg/grpcserver/v1"
	"github.com/aprln/trading-chart/websocket-listener/config"
	"github.com/aprln/trading-chart/websocket-listener/internal/handler"
	"github.com/aprln/trading-chart/websocket-listener/internal/service"
	binanceConnector "github.com/binance/binance-connector-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		<-interrupt
		os.Exit(0)
	}()

	cfg := config.New()

	wsStreamClient := binanceConnector.NewWebsocketStreamClient(false, cfg.BinanceWebsocketURL)
	if wsStreamClient == nil {
		log.Fatal("Cannot init WS client")

		return
	}



	var wg sync.WaitGroup
	wg.Add(len(cfg.TradeSymbols))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, symbol := range cfg.TradeSymbols {
		go processSymbol(ctx, symbol, wsStreamClient, cfg.GRPCServerURL, &wg)
	}

	wg.Wait()
}

func processSymbol(
	ctx context.Context,
	symbol string,
	wsStreamClient *binanceConnector.WebsocketStreamClient,
	grpcServerURL string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	fmt.Println("SYMBOL: ", symbol)

	conn, err := grpc.NewClient(grpcServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Create one handler per symbol
	hdl := handler.New(
		service.New(
			v1.NewTradingChartClient(conn),
		),
	)

	doneCh, _, err := wsStreamClient.WsAggTradeServe(
		symbol,
		func(event *binanceConnector.WsAggTradeEvent) {
			hdl.HandleAggTrade(ctx, event)
		},
		handleError,
	)
	if err != nil {
		log.Println(err)

		return
	}

	<-doneCh

	return
}

func handleError(err error) {
	log.Println(err)
}
