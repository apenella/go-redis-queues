package stream

import (
	"context"
	"fmt"
	"math/rand"
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
func (c *Consumer) Consume(ctx context.Context, stream string) {
	defer c.client.Close()

	result := make(chan interface{})

	go func() {

		clientId := string(rand.Int())

		// xreadargs := &redis.XReadArgs{
		// 	Streams: []string{stream, "$"},
		// 	Block:   0,
		// }
		xreadgroupargs := &redis.XReadGroupArgs{
			Group:    "default",
			Consumer: clientId,
			Streams:  []string{stream, ">"},
		}

		for {

			message, err := c.client.XReadGroup(xreadgroupargs).Result()
			if err != nil {
				fmt.Println("Error reading stream", stream, err.Error())
			} else {
				result <- message
			}
		}
	}()

	for {
		select {
		case m := <-result:
			fmt.Println(m)
		case <-time.After(5 * time.Second):
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
