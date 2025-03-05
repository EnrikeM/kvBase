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

var (
	errEmptyQuery = fmt.Errorf("empty query")
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(query string) (Query, error) {
	querySplit := strings.Split(query, " ")
	if len(querySplit) == 0 {
		return Query{}, errEmptyQuery
	}

	switch Method(querySplit[0]) {
	case SET:
		return handleSET(querySplit)
	case GET:
		return handleGET(querySplit)
	case DEL:
		return handleDEL(querySplit)
	}

	return Query{}, fmt.Errorf("command %s is unsupported", querySplit[0])
}
func handleSET(querySplit []string) (Query, error) {
	if len(querySplit) != 3 {
		return Query{}, fmt.Errorf("invalid set command")
	}

	return Query{
		Method: SET,
		Args:   []string{querySplit[1], querySplit[2]},
	}, nil
}

func handleGET(querySplit []string) (Query, error) {
	if len(querySplit) != 2 {
		return Query{}, fmt.Errorf("invalid get command")
	}

	return Query{
		Method: GET,
		Args:   []string{querySplit[1]},
	}, nil
}

func handleDEL(querySplit []string) (Query, error) {
	if len(querySplit) != 2 {
		return Query{}, fmt.Errorf("invalid del command")
	}

	return Query{
		Method: SET,
		Args:   []string{querySplit[1]},
	}, nil
}
