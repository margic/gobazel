package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/margic/gobazel/protos"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	var addr string
	flag.StringVarP(&addr, "listen", "l", ":8081", "address greeting service is listening on")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("Starting greeting service")

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
	protos.RegisterGreetingServer(grpcServer, &greetingServer{
		l: logger,
	})
	grpcServer.Serve(lis)
}

type greetingServer struct {
	l *zap.Logger
}

func (gs *greetingServer) Greeting(ctx context.Context, req *protos.GreetingRequest) (*protos.GreetingResponse, error) {
	res := &protos.GreetingResponse{
		Greeting:      "Hello " + req.Name,
		ServerTime:    time.Now().String(),
		MessageID:     newMessageID(),
		CorrelationID: req.MessageID,
		Hostname:      os.Getenv("HOSTNAME"),
		Err:           "",
	}
	gs.l.Info("Greeting Response",
		zap.String("name", req.Name),
		zap.String("greeting", res.Greeting),
		zap.String("messageID", res.MessageID),
		zap.String("correlationID", res.CorrelationID),
	)
	return res, nil
}

func newMessageID() string {
	id := uuid.NewV1()
	return id.String()
}
