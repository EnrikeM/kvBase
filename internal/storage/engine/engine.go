package engine

import (
	"fmt"
)

type Engine struct {
	memory HashTable
}

func NewEngine() *Engine {
	return &Engine{
		memory: HashTable{
			data: make(map[string]string),
		},
	}
}

type HashTable struct {
	data map[string]string
}

func (s *Engine) Set(args []string) (string, error) {
	s.memory.data[args[0]] = args[1]
	return fmt.Sprintf("key %s set for val %s", args[0], args[1]), nil
}

func (s *Engine) Get(args []string) (string, error) {
	val := s.memory.data[args[0]]
	return val, nil
}

func (s *Engine) Del(args []string) (string, error) {
	delete(s.memory.data, args[0])
	return fmt.Sprintf("%s deleted", args[0]), nil
}
