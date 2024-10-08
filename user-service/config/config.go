package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	ErrOnLoad    = errors.New("failed to load env from file")
	ErrOnProcess = errors.New("failed to process env")
)

type Config struct {
	PgDSN        string `envconfig:"PG_DSN" required:"true"`
	RABBITMQ_URL string `envconfig:"RABBITMQ_URL" required:"true"`
	Port         string `envconfig:"PORT" default:":8080"`
	GrpcPort     string `envconfig:"GRPC_PORT" default:":50001"`
}

func Load(filenames ...string) (Config, error) {
	// by default loads from .env
	err := godotenv.Load(filenames...)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.Join(ErrOnLoad, err)
	}

	conf := Config{}
	err = envconfig.Process("", &conf)
	if err != nil {
		return Config{}, errors.Join(ErrOnProcess, err)
	}

	return conf, nil
}
