package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := greetingService{}

	greetingHandler := httptransport.NewServer(
		makeGreetingEndpoint(svc),
		decodeGreetingRequest,
		encodeResponse,
	)

	http.Handle("/", greetingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeGreetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := greetingRequest{}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// GreetingService returns a greeting
type GreetingService interface {
	Greeting(context.Context) (string, error)
}

// greetingService a type that implementes GreetingService
type greetingService struct{}

func (greetingService) Greeting(_ context.Context) (string, error) {
	return "service2 " + time.Now().String(), nil
}

// greetingRequest is empty as there are no params for request but type required
type greetingRequest struct {
}

// greetingResponse contains the response
type greetingResponse struct {
	Greeting string `json:"greeting"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// makeGreetingEndpoint wrap service with endpoint
func makeGreetingEndpoint(svc GreetingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(greetingRequest) ignore request but it would normally be here
		v, err := svc.Greeting(ctx)
		if err != nil {
			return greetingResponse{v, err.Error()}, nil
		}
		return greetingResponse{v, ""}, nil
	}
}
