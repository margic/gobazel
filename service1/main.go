package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// a call to viper to test build with viper
	viper.AutomaticEnv()
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("Started After Update")
}
