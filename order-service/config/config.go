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
	MongoURL      string `envconfig:"MONGO_URL" required:"true"`
	Port          string `envconfig:"PORT" default:":8081"`
	MongoUsername string `envconfig:"MONGO_USERNAME" required:"true"`
	MongoPassword string `envconfig:"MONGO_PASSWORD" required:"true"`
	RABBITMQ_URL  string `envconfig:"RABBITMQ_URL" required:"true"`
	GRPC_URL      string `envconfig:"GRPC_URL" required:"true"`
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
