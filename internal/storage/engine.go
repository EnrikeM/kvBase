package storage

import (
	"fmt"

	"github.com/EnrikeM/kvBase/internal/compute"
)

const (
	errWrongNum    = "wrong number of arguments for method %s, expected %d, got %d"
	errNoSuchField = "no such field %s"
)

type EngineSrvc struct {
	memory HashTable
}

func NewEngineSrvc() *EngineSrvc {
	return &EngineSrvc{
		memory: HashTable{
			data: make(map[string]string, 0),
		},
	}
}

type HashTable struct {
	data map[string]string
}

func (s *EngineSrvc) Set(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf(errWrongNum, compute.SET, 2, len(args))
	}
	s.memory.data[args[0]] = args[1]

	return fmt.Sprintf("key %s set for val %s", args[0], args[1]), nil
}

func (s *EngineSrvc) Get(args []string) (string, error) {
	val, ok := s.memory.data[args[0]]
	if !ok {
		return "", fmt.Errorf(errNoSuchField, args[0])
	}

	return val, nil
}

func (s *EngineSrvc) Del(args []string) (string, error) {
	_, ok := s.memory.data[args[0]]
	if !ok {
		return "", fmt.Errorf(errNoSuchField, args[0])
	}

	delete(s.memory.data, args[0])

	return fmt.Sprintf("%s deleted", args[0]), nil
}
