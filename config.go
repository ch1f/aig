package aig

import (
	"os"

	"go.mongodb.org/mongo-driver/event"
)

const (
	envURI        = "AIG_MONGO_URI"
	envUser       = "AIG_MONGO_USER"
	envPassword   = "AIG_MONGO_PASSWORD"
	envDB         = "AIG_MONGO_DB"
	envCollection = "AIG_MONGO_COLLECTION"
)

type MongoConfig struct {
	URI        string `env:"AIG_MONGO_URI" envDefault:"mongodb://127.0.0.1:27017"`
	User       string `env:"AIG_MONGO_USER"`
	Password   string `env:"AIG_MONGO_PASSWORD"`
	DB         string `env:"AIG_MONGO_DB" envDefault:"generator"`  //example: generator
	Collection string `env:"AIG_MONGO_COLLECTION" envDefault:"ai"` //example: ai
	Monitor    *event.CommandMonitor
}

func GetConfigEnv() MongoConfig {
	cfg := DefaultConfig()

	envString(envURI, &cfg.URI)
	envString(envUser, &cfg.User)
	envString(envPassword, &cfg.Password)
	envString(envDB, &cfg.DB)
	envString(envCollection, &cfg.Collection)

	return cfg
}

func DefaultConfig() MongoConfig {
	return MongoConfig{
		URI:        "mongodb://127.0.0.1:27017",
		DB:         "generator",
		Collection: "ai",
	}
}

func envString(env string, v *string) {
	if val, ok := os.LookupEnv(env); ok {
		*v = val
	}
}
