package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/margic/gobazel/protos"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	var addr string
	flag.StringVarP(&addr, "listen", "l", ":8080", "address greeting service is listening on")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("starting greeting service")

	myCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Counter",
		})

	prometheus.MustRegister(myCounter)

	go func(logger *zap.Logger) {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9090", nil)
	}(logger)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to start listener",
			zap.String("error", err.Error()),
			zap.String("address", addr))
	}

	grpcServer := grpc.NewServer()
	protos.RegisterGreetingServer(grpcServer, &greetingServer{})
	grpcServer.Serve(lis)
}

type greetingServer struct {
}

func (gs *greetingServer) Greeting(ctx context.Context, _ *protos.Empty) (*protos.GreetingResponse, error) {
	return &protos.GreetingResponse{
		Greeting: "Hello world at: " + time.Now().String(),
		Err:      "",
	}, nil
}
