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
	cfg := GetConfigEnv()
	aiM := Create(context.Background(), cfg)

	s.repo = aiM
}

func TestMongo(t *testing.T) {
	suite.Run(t, new(Suite))
}
