package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoDB ConfigMongo
	Server  ConfigServer
}
type ConfigMongo struct {
	URI      string
	UserName string
	Password string
	Database string
}

type ConfigServer struct {
	Port int
}

func NewCfg() (*Config, error) {
	if err := godotenv.Load(); err != nil {

		return nil, err
	}
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.MongoDB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil

}
