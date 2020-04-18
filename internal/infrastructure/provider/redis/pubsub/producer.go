package pubsub

import (
	"context"

	"github.com/go-redis/redis/v7"

	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	providerredis "github.com/apenella/go-redis-queues/internal/infrastructure/provider/redis"
)

// Producer is a redis client which appends messages to a channel
type Producer struct {
	client *redis.Client
}

// NewClient creates a new redis client
func NewProducer(config *configuration.Configuration) (*Producer, error) {
	c, err := providerredis.NewClient(config.RedisHost, config.RedisPort, config.RedisDB, config.RedisPassword)
	if err != nil {
		return nil, err
	}

	producer := &Producer{
		client: c,
	}

	return producer, nil
}

// Publish appends a new event to channel
func (p *Producer) Publish(ctx context.Context, channel string, messages interface{}) {
	p.client.Publish(channel, messages)
}

// Publish appends a new event to channel
func (p *Producer) Ping(ctx context.Context) error {
	_, err := p.client.Ping().Result()
	if err != nil {
		return err
	}

	p.client.Close()

	return nil
}
