package compute

import (
	"fmt"
	"strings"
)

type Query struct {
	Method Method
	Args   []string
}

type Method string

var (
	SET Method = "SET"
	GET Method = "GET"
	DEL Method = "DEL"
)

const (
	errEmptyQuery = ("empty query %s")
	errBadMethod  = ("method %s unsupported")
	errWrongNum   = "wrong number of arguments for method %s, expected %d, got %d"
)

type ParserSrvc struct {
}

func NewParserSrvc() *ParserSrvc {
	return &ParserSrvc{}
}

func (p *ParserSrvc) Parse(query string) (Query, error) {
	querySplit := strings.Split(query, " ")
	if len(querySplit) == 0 {
		return Query{}, fmt.Errorf(errEmptyQuery, query)
	}

	switch Method(querySplit[0]) {
	case SET:
		return handleSET(querySplit)
	case GET:
		return handleGET(querySplit)
	case DEL:
		return handleDEL(querySplit)
	}

	return Query{}, fmt.Errorf(errBadMethod, querySplit[0])
}
func handleSET(querySplit []string) (Query, error) {
	if len(querySplit) != 3 {
		return Query{}, fmt.Errorf(errWrongNum, SET, 3, len(querySplit))
	}

	return Query{
		Method: SET,
		Args:   []string{querySplit[1], querySplit[2]},
	}, nil
}

func handleGET(querySplit []string) (Query, error) {
	if len(querySplit) != 2 {
		return Query{}, fmt.Errorf(errWrongNum, GET, 2, len(querySplit))
	}

	return Query{
		Method: GET,
		Args:   []string{querySplit[1]},
	}, nil
}

func handleDEL(querySplit []string) (Query, error) {
	if len(querySplit) != 2 {
		return Query{}, fmt.Errorf(errWrongNum, DEL, 2, len(querySplit))
	}

	return Query{
		Method: DEL,
		Args:   []string{querySplit[1]},
	}, nil
}
