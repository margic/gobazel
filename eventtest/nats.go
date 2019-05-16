package main

import (
	"github.com/nats-io/nats"
	"github.com/nats-io/nats/encoders/protobuf"
	"go.uber.org/zap"
)

// NatsClient is a wrapper around a nats connection with logging
type NatsClient struct {
	logger *zap.Logger
	nc     *nats.EncodedConn
}

// NatsConfig type to store values required to create nats client
type NatsConfig struct {
	Logger *zap.Logger
	URL    string
}

// Close function to allow deferred close of nats connection
func (n *NatsClient) Close() {
	n.nc.Close()
}

// NewNatsClient creates a new nats client
func NewNatsClient(cfg *NatsConfig) (*NatsClient, error) {
	nc, err := nats.Connect(cfg.URL)

	if err != nil {
		return nil, err
	}
	nats.RegisterEncoder(protobuf.PROTOBUF_ENCODER, &protobuf.ProtobufEncoder{})
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}
	client := &NatsClient{
		logger: cfg.Logger,
		nc:     c,
	}
	return client, nil
}
