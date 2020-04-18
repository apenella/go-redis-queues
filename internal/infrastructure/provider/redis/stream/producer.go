package stream

import (
	"context"
	"fmt"

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
func (p *Producer) Publish(ctx context.Context, stream string, message interface{}) {

	xaddargs := &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{
			"message": message,
		},
	}

	xgroup := "default"

	if !existsXGroup(p.client, xgroup) {
		status := p.client.XGroupCreateMkStream(stream, "default", "$")
		if status.Err() != nil {
			fmt.Println(status.Err().Error())
			return
		}
	}

	p.client.XAdd(xaddargs)
}

// Publish appends a new event to channel
func (p *Producer) Ping(ctx context.Context) error {
	_, err := p.client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
