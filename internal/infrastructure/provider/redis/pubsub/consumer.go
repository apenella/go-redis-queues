package pubsub

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	providerredis "github.com/apenella/go-redis-queues/internal/infrastructure/provider/redis"
)

// Producer is a redis client which appends messages to a channel
type Consumer struct {
	client *redis.Client
}

// NewClient creates a new redis client
func NewConsumer(config *configuration.Configuration) (*Consumer, error) {
	c, err := providerredis.NewClient(config.RedisHost, config.RedisPort, config.RedisDB, config.RedisPassword)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		client: c,
	}

	return consumer, nil
}

// Publish appends a new event to channel
func (c *Consumer) Consume(ctx context.Context, channel string) {
	pubsub := c.client.Subscribe(channel)
	defer pubsub.Close()

	subChannel := pubsub.Channel()

	for {
		select {
		case m := <-subChannel:
			fmt.Println(m.String())
		case <-time.After(5 * time.Second):
			fmt.Println("Consumer cancelled after 5 idle seconds")
			c.client.Close()
			return
		case <-ctx.Done():
			fmt.Println("Consumer cancelled by user")
			c.client.Close()
			return
		}
	}
}

// Publish appends a new event to channel
func (c *Consumer) Ping(ctx context.Context) error {
	_, err := c.client.Ping().Result()
	if err != nil {
		return err
	}

	c.client.Close()

	return nil
}
