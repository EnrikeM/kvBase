package main

import (
	"log"

	"github.com/EnrikeM/kvBase/internal/compute"
	"github.com/EnrikeM/kvBase/internal/compute/parser"
	"github.com/EnrikeM/kvBase/internal/storage"
	"github.com/EnrikeM/kvBase/internal/storage/engine"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	strg := storage.NewService(
		engine.NewEngine(),
		compute.NewService(logger, parser.NewParser()),
		logger,
	)

	res, _ := strg.Update("SET Hero Cheburashka")
	log.Println(res)
	res, _ = strg.Update("GET Hero")
	log.Println(res)
	res, _ = strg.Update("DEL Hero")
	log.Println(res)

	res, _ = strg.Update("GET Hero")
	log.Println(res)
}
