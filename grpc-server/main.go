package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/aprln/trading-chart/grpc-server/config"
	"github.com/aprln/trading-chart/grpc-server/internal/handler"
	"github.com/aprln/trading-chart/grpc-server/internal/repository"
	"github.com/aprln/trading-chart/grpc-server/internal/service"
	grpcServerV1 "github.com/aprln/trading-chart/grpc-server/pkg/grpcserver/v1"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const dbDriverPostgres = "postgres"

func main() {
	cfg := config.New()

	db, err := connectDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalln(err)
	}

	if err = runGRPCServer(
		cfg.GRPCServerPort,
		handler.New(
			service.New(
				repository.New(db),
			),
		),
	); err != nil {
		log.Fatalf("failed to create spanner client %v", err)
	}
}

func runGRPCServer(port uint32, hdl *handler.Handler) error {
	netListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcMiddleware.ChainStreamServer(
				grpcRecovery.StreamServerInterceptor(), // TODO: Add error interceptor
			),
		),
	)

	grpcServerV1.RegisterTradingChartServer(grpcServer, hdl)

	log.Println("Running server... ")
	if err := grpcServer.Serve(netListener); err != nil {
		return err
	}

	return nil
}

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriverPostgres, dsn)
	if err != nil {
		return nil, err
	}

	// Ensure connection is valid with a ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
