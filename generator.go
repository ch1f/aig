package aig

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *AI) Next(name string) uint64 {
	ctx := context.Background()
	filter := m.getFilter(name)
	update := m.getUpdate()
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	result := m.db.Collection(m.collectionName).FindOneAndUpdate(ctx, filter, update, opts)

	err := result.Err()
	if err != nil {
		log.Println("ai next do error:", errors.WithStack(err))
		return 0
	}

	value := counter{}
	if errDecode := result.Decode(&value); errDecode != nil {
		log.Println("ai next decode error", errors.WithStack(errDecode))
		return 0
	}

	return value.Seq
}

func (m *AI) Generate(ctx context.Context, name string) uint64 {
	filter := m.getFilter(name)
	update := m.getUpdate()
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	result := m.db.Collection(m.collectionName).FindOneAndUpdate(ctx, filter, update, opts)

	err := result.Err()
	if err != nil {
		log.Println("ai next do error:", errors.WithStack(err))
		return 0
	}

	value := counter{}
	if errDecode := result.Decode(&value); errDecode != nil {
		log.Println("ai next decode error", errors.WithStack(errDecode))
		return 0
	}

	return value.Seq
}
