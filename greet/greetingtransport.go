package main

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/margic/gobazel/protos"
)

// NewGRPCClient returns an AddService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func makeGreetingEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	// Each individual endpoint is an http/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares.
	var greetingEndpoint endpoint.Endpoint

	greetingEndpoint = grpctransport.NewClient(
		conn,
		"greeting",
		"greeting",
		encodeGRPCGreetingRequest,
		decodeGRPCGreetingResponse,
		protos.GreetingResponse{},
	).Endpoint()

	return greetingEndpoint
}

func decodeGRPCGreetingResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeGRPCGreetingRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}
