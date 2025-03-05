package storage

import (
	"fmt"

	"github.com/EnrikeM/kvBase/internal/compute"
	"go.uber.org/zap"
)

//go:generate mockgen -source=storage.go -destination=storage_mocks_test.go -package=storage_test

type Service struct {
	Compute
	Engine
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {
	return &Service{
		Engine:  NewEngineSrvc(),
		Compute: compute.NewService(logger, *compute.NewParser()),
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
	compQuery, err := s.HandleQuery(query)
	if err != nil {
		s.logger.Error("err handling query", zap.Error(err))
		return "", err
	}
	switch compQuery.Method {
	case compute.SET:
		return s.Set(compQuery.Args)
	case compute.GET:
		return s.Get(compQuery.Args)
	case compute.DEL:
		return s.Del(compQuery.Args)
	}

	s.logger.Error("err processing query", zap.String("query", query))
	return "", fmt.Errorf("unknown method")
}
