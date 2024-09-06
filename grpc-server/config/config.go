package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env            string
	GRPCServerPort uint32
	DatabaseDSN    string
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
	viper.AddConfigPath("./grpc-server")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg = &Config{
		Env:            viper.GetString("ENV"),
		GRPCServerPort: viper.GetUint32("GRPC_SERVER_PORT"),
		DatabaseDSN:    viper.GetString("DATABASE_DSN"),
	}
}
