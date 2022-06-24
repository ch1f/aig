package aig

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	repo *AI
}

func (s *Suite) SetupSuite() {
	aiM := Create(context.Background(), &MongoConfig{
		URI:        "mongodb://127.0.0.1:27017",
		DB:         "testAI",
		Collection: "ai",
	})

	s.repo = aiM
}

func TestMongo(t *testing.T) {
	suite.Run(t, new(Suite))
}
