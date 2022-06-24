package aig

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type MongoConfig struct {
	URI        string `env:"AIG_MONGO_URI,required"` //example: mongodb://127.0.0.1:27017
	User       string `env:"AIG_MONGO_USER"`
	Password   string `env:"AIG_MONGO_PASSWORD"`
	DB         string `env:"AIG_MONGO_DB,required"`         //example: generator
	Collection string `env:"AIG_MONGO_COLLECTION,required"` //example: ai
}

func GetConfigEnv() (*MongoConfig, error) {
	cfg := MongoConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	return &cfg, nil
}
