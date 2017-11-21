package main

import (
	"flag"
	"fmt"
	"net/http"
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
	logger.Info("starting service 1")
	http.HandleFunc("/", HandleHello)
	logger.Fatal(http.ListenAndServe(":8080", nil).Error())
}

// HandleHello send hello to client
func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}
