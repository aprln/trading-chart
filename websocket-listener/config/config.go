package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env                 string
	GRPCServerURL       string
	BinanceWebsocketURL string
	TradeSymbols        []string
}

var cfg *Config

func New() *Config {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./websocket-listener")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg = &Config{
		Env:                 viper.GetString("ENV"),
		GRPCServerURL:       viper.GetString("GRPC_SERVER_URL"),
		BinanceWebsocketURL: viper.GetString("BINANCE_WEBSOCKET_URL"),
		TradeSymbols:        strings.Split(viper.GetString("TRADE_SYMBOLS"), "|"),
	}
}
