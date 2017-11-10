package main

import (
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("Started")
}
