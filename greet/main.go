package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	var addr string
	flag.StringVarP(&addr, "greetingAddr", "g", "greeting.gobazel", "Address of Greeting Service, defaults to greeting.gobazel")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Unable to create logger: %s\n", err)
		os.Exit(1)
	}
	logger.Info("starting greet service")

	// clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	// if err != nil {
	// 	logger.Fatal("error creating connection to greeting services",
	// 		zap.String("error", err.Error()),
	// 		zap.String("address", addr))
	// }

	// svc := greet{
	// 	greeting: greetingProxy{
	// 		greeting: makeGreetingEndpoint(clientConn),
	// 	},
	// }

	//http.Handle("/", greetHandler)
	logger.Fatal(http.ListenAndServe(":8080", nil).Error())
}

// func decodegreetRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	request := greetRequest{}
// 	return request, nil
// }

// func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// 	return json.NewEncoder(w).Encode(response)
// }

// // greeter returns a greet
// type greeter interface {
// 	greet(context.Context) (string, error)
// }

// // greet a type that implementes greetService
// type greet struct {
// 	greeting greetingProxy
// }

// func (g greet) greet(ctx context.Context) (string, error) {
// 	resp, err := g.greeting.Greeting(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp, nil
// }

// // greetRequest is empty as there are no params for request but type required
// type greetRequest struct {
// }

// // greetResponse contains the response
// type greetResponse struct {
// 	Greet string `json:"greet"`
// 	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
// }

// // makegreetEndpoint wrap service with endpoint
// func makegreetEndpoint(svc greet) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		//req := request.(greetRequest) ignore request but it would normally be here
// 		v, err := svc.greet(ctx)
// 		if err != nil {
// 			return greetResponse{v, err.Error()}, nil
// 		}
// 		return greetResponse{v, ""}, nil
// 	}
// }

// // Add cliet call to greeting service
// // greeting proxy is the proxy around the greeting service we will call
// type greetingProxy struct {
// 	greeting endpoint.Endpoint
// }

// func (gp greetingProxy) Greeting(ctx context.Context) (string, error) {
// 	// really all we are doing here is unwrapping the grpc response values
// 	response, err := gp.greeting(ctx, protos.Empty{})
// 	if err != nil {
// 		return "", err
// 	}
// 	resp := response.(protos.GreetingResponse)
// 	if resp.Err != "" {
// 		return "", errors.New(resp.Err)
// 	}
// 	return "", nil
// }
