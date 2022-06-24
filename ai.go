package aig

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type AI struct {
	db             *mongo.Database
	collectionName string
}

type counter struct {
	Seq uint64 `bson:"seq"`
}

const timeout = time.Second * 10

func Create(ctx context.Context, cfg *MongoConfig, opts ...*SettingOptions) *AI {
	clientOpts := options.Client().ApplyURI(cfg.URI).
		SetWriteConcern(writeconcern.New(
			writeconcern.WTimeout(timeout),
			writeconcern.J(false),
		)).SetRetryWrites(false)

	//set options
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.Otel != nil {
			clientOpts.Monitor = otelmongo.NewMonitor()
		}
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
