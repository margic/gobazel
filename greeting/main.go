package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var myCounter prometheus.Counter

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

	myCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Counter",
		})

	prometheus.MustRegister(myCounter)

	go func(logger *zap.Logger) {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9090", nil)
	}(logger)

	http.HandleFunc("/", HandleHello)
	logger.Fatal(http.ListenAndServe(":8080", nil).Error())
}

// HandleHello send hello to client
func HandleHello(w http.ResponseWriter, r *http.Request) {
	defer myCounter.Add(1)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting()))
}
