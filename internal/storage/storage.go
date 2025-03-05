package storage

import (
	"fmt"

	"github.com/EnrikeM/kvBase/internal/compute"
	"go.uber.org/zap"
)

//go:generate mockgen -source=storage.go -destination=storage_mocks_test.go -package=storage_test

type Service struct {
	compute Compute
	engine  Engine
	logger  *zap.Logger
}

func NewService(engine Engine, compute Compute, logger *zap.Logger) *Service {
	return &Service{
		logger:  logger,
		engine:  engine,
		compute: compute,
	}
}

type Compute interface {
	HandleQuery(query string) (compute.Query, error)
}

type Engine interface {
	Set(args []string) (string, error)
	Get(args []string) (string, error)
	Del(args []string) (string, error)
}

func (s *Service) Update(query string) (string, error) {
	compQuery, err := s.compute.HandleQuery(query)
	if err != nil {
		s.logger.Error("err handling query", zap.Error(err))
		return "", err
	}
	switch compQuery.Method {
	case compute.SET:
		return s.engine.Set(compQuery.Args)
	case compute.GET:
		return s.engine.Get(compQuery.Args)
	case compute.DEL:
		return s.engine.Del(compQuery.Args)
	}

	s.logger.Error("err processing query", zap.String("query", query))
	return "", fmt.Errorf("unknown method")
}
