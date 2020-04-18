package transport

import (
	"context"
	"errors"

	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/internal/infrastructure/provider/redis/fifo"
	"github.com/apenella/go-redis-queues/internal/infrastructure/provider/redis/pubsub"
	"github.com/apenella/go-redis-queues/internal/infrastructure/provider/redis/stream"
)

type Producer interface {
	Publish(ctx context.Context, channel string, messages interface{})
}

type Consumer interface {
	Consume(ctx context.Context, channel string)
}

type Transport struct {
	producer Producer
	comsumer Consumer
}

const (
	newTransportUnknownType          = "Unkown transport type"
	newTransportUnavailableTransport = "Transport type unavailable"
)

func NewTransport(config *configuration.Configuration) (*Transport, error) {
	var err error
	transport := &Transport{}

	switch config.Transport {
	case configuration.TransportPubsub:
		transport.producer, err = pubsub.NewProducer(config)
		if err != nil {
			return nil, err
		}
		transport.comsumer, err = pubsub.NewConsumer(config)
		if err != nil {
			return nil, err
		}

	case configuration.TransportStream:
		transport.producer, err = stream.NewProducer(config)
		if err != nil {
			return nil, err
		}
		transport.comsumer, err = stream.NewConsumer(config)
		if err != nil {
			return nil, err
		}

	case configuration.TransportFifo:
		transport.producer, err = fifo.NewProducer(config)
		if err != nil {
			return nil, err
		}
		transport.comsumer, err = fifo.NewConsumer(config)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(newTransportUnknownType)
	}

	return transport, nil
}

func (c *Transport) Publish(ctx context.Context, channel string, messages interface{}) {
	c.producer.Publish(ctx, channel, messages)
}

func (c *Transport) Consume(ctx context.Context, channel string) {
	c.comsumer.Consume(ctx, channel)
}
