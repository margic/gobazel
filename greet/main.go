package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/margic/gobazel/protos"
	"google.golang.org/grpc"

	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	var addr string
	flag.StringVarP(&addr, "greetingAddr", "g", "greeting.gobazel", "Address of Greeting Service, defaults to greeting.gobazel")
	flag.Parse()

	logger, err := zap.NewProduction(zap.Fields(zap.String("Service", "Greet")))

	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("starting greet service")

	http.HandleFunc("/", handler(logger))
	logger.Fatal(http.ListenAndServe(":8080", nil).Error())
}

func handler(l *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial("greeting.gobazel:80", grpc.WithInsecure())
		if err != nil {
			l.Error("Error connected to greeting service", zap.Error(err))
		}
		defer conn.Close()
		c := protos.NewGreetingClient(conn)
		g, err := c.Greeting(r.Context(), &protos.Empty{})
		if err != nil {
			l.Error("error calling grpc", zap.Error(err))
		}
		greeting := &greetResponse{
			Greet:    g.Greeting,
			Hostname: os.Getenv("HOSTNAME"),
		}
		enc := json.NewEncoder(w)
		w.Header().Add("Content-Type", "application/json")
		enc.Encode(greeting)
		l.Info("Handler returned response", zap.String("greet", greeting.Greet))
	}
}

// greetRequest is empty as there are no params for request but type required
type greetRequest struct {
}

// greetResponse contains the response
type greetResponse struct {
	Greet    string `json:"greet"`
	Hostname string `json:"hostname"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
