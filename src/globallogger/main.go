package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	// Default Global Logger
	fmt.Printf("\n*** Using the default global logger\n\n")

	zap.S().Infow("An info message", "iteration", 1)
	zap.S().Sync()

	fmt.Printf("\n*** After replacing the global logger with a development logger\n\n")

	// Development Global Logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)
	zap.S().Sync()

	// Production Global Logger
	fmt.Printf("\n*** After replacing the global logger with a production logger\n\n")
	logger, _ = zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)
	zap.S().Sync()

	fmt.Printf("\n*** After undoing the last replacement of the global logger\n\n")
	undo()
	zap.S().Infow("An info message", "iteration", 1)
	zap.S().Sync()
}
