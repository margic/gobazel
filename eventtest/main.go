package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	var addr string
	var wait time.Duration
	flag.StringVarP(&addr, "natsAddr", "n", "nats.gobazel", "Address of Nats Service, defaults to nats.gobazel")
	flag.DurationVar(&wait, "shutdown-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logger, err := zap.NewProduction(zap.Fields(zap.String("Service", "EventTest")))

	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}

	// create aa new handlers instance with logger
	handlers := Handlers{
		logger: logger,
	}

	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/health", handlers.HealthHandler)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Info("Starting EventTest service", zap.String("natsAddr", addr))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Info("Stopping EventTest service")
	os.Exit(0)
}
