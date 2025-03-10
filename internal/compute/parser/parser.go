package parser

import (
	"fmt"
	"strings"

	"github.com/EnrikeM/kvBase/internal/domain"
)

const (
	errEmptyQuery = ("empty query %s")
	errBadMethod  = ("method %s unsupported")
	errWrongNum   = "wrong number of arguments for method %s, expected %d, got %d"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(query string) (domain.Query, error) {
	querySplit := strings.Split(query, " ")
	if len(querySplit) == 0 {
		return domain.Query{}, fmt.Errorf(errEmptyQuery, query)
	}

	switch domain.Method(querySplit[0]) {
	case domain.SET:
		return handleSET(querySplit)
	case domain.GET:
		return handleGET(querySplit)
	case domain.DEL:
		return handleDEL(querySplit)
	}

	return domain.Query{}, fmt.Errorf(errBadMethod, querySplit[0])
}
func handleSET(querySplit []string) (domain.Query, error) {
	if len(querySplit) != domain.SetLen {
		return domain.Query{}, fmt.Errorf(errWrongNum, domain.SET, domain.SetLen, len(querySplit))
	}

	return domain.Query{
		Method: domain.SET,
		Args:   []string{querySplit[1], querySplit[2]},
	}, nil
}

func handleGET(querySplit []string) (domain.Query, error) {
	if len(querySplit) != domain.GetLen {
		return domain.Query{}, fmt.Errorf(errWrongNum, domain.GET, domain.SetLen, len(querySplit))
	}

	return domain.Query{
		Method: domain.GET,
		Args:   []string{querySplit[1]},
	}, nil
}

func handleDEL(querySplit []string) (domain.Query, error) {
	if len(querySplit) != domain.DelLen {
		return domain.Query{}, fmt.Errorf(errWrongNum, domain.DEL, domain.DelLen, len(querySplit))
	}

	return domain.Query{
		Method: domain.DEL,
		Args:   []string{querySplit[1]},
	}, nil
}
