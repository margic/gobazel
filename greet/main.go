package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"

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
	logger.Info("Starting greet service")

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

		name, ok := r.URL.Query()["name"]
		if !ok || len(name) == 0 {
			name = []string{"Stranger"}
		}

		// build request
		gr := &protos.GreetingRequest{
			MessageID: newMessageID(),
			Name:      name[0],
		}

		// send request
		g, err := c.Greeting(r.Context(), gr)
		if err != nil {
			l.Error("error calling grpc", zap.Error(err))
		}

		// extract greeting from response
		greeting := &greetResponse{
			Greet:         g.Greeting,
			GreetHost:     os.Getenv("HOSTNAME"),
			GreetingHost:  g.Hostname,
			OrgMessageID:  gr.MessageID,
			ResMessageID:  g.MessageID,
			CorrelationID: g.CorrelationID,
			ServerTime:    g.ServerTime,
		}
		enc := json.NewEncoder(w)
		w.Header().Add("Content-Type", "application/json")
		enc.Encode(greeting)
		l.Info("Greet response",
			zap.String("greet", greeting.Greet),
			zap.String("greethostname", greeting.GreetHost),
			zap.String("greetinghostname", greeting.GreetingHost),
			zap.String("correlationID", greeting.CorrelationID),
			zap.String("servertime", g.ServerTime),
		)
	}
}

func newMessageID() string {
	id := uuid.NewV1()
	return id.String()
}

// greetRequest is empty as there are no params for request but type required
type greetRequest struct {
}

// greetResponse contains the response
type greetResponse struct {
	Greet         string `json:"greet"`
	GreetHost     string `json:"greethost"`
	GreetingHost  string `json:"greetinghost"`
	OrgMessageID  string `json:"orgmessageid"`
	ResMessageID  string `json:"resmessageid"`
	CorrelationID string `json:"correlationid"`
	ServerTime    string `json:"servertime"`
	Err           string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
