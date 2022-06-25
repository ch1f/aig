package aig

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type AI struct {
	db             *mongo.Database
	collectionName string
}

type counter struct {
	Seq uint64 `bson:"seq"`
}

const timeout = time.Second * 10

func Create(ctx context.Context, cfg MongoConfig) *AI {
	clientOpts := options.Client().ApplyURI(cfg.URI).
		SetWriteConcern(writeconcern.New(
			writeconcern.WTimeout(timeout),
			writeconcern.J(false),
		)).SetRetryWrites(false)

	if cfg.User != "" && cfg.Password != "" {
		clientOpts.SetAuth(options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		})
	}

	if cfg.Monitor != nil {
		clientOpts.Monitor = cfg.Monitor
	}

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic("ai error connect to mongoDB" + err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic("ai error ping to mongoDB" + err.Error())
	}

	ai := AI{
		db:             client.Database(cfg.DB),
		collectionName: cfg.Collection,
	}
	return &ai
}
