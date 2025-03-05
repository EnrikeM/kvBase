package compute

import (
	"fmt"

	"go.uber.org/zap"
)

const (
	errBadQuery = ("err parsing query %s")
)

//go:generate mockgen -source=compute.go -destination=compute_mocks_test.go -package=compute_test

type Service struct {
	logger *zap.Logger
	parser Parser
}

func NewService(logger *zap.Logger, parser Parser) *Service {
	return &Service{
		logger: logger,
		parser: parser,
	}
}

func (s *Service) HandleQuery(query string) (Query, error) {
	parsedQuery, err := s.parser.Parse(query)
	if err != nil {
		s.logger.Error("err parsing query", zap.String("query", query))
		return Query{}, fmt.Errorf(errBadQuery, query)
	}

	return parsedQuery, nil
}
