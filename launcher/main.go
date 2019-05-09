package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

func main() {
	// var addr string
	// flag.StringVarP(&addr, "greetingAddr", "g", "greeting.gobazel", "Address of Greeting Service, defaults to greeting.gobazel")
	// flag.Parse()

	logger, err := zap.NewProduction(zap.Fields(zap.String("Service", "Greet")))

	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	var srv http.Server

	//idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Info("HTTP server Shutdown: %v", zap.Error(err))
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	logger.Info("Starting launcher service")
}
