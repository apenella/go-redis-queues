package redis

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis/v7"
)

// NewClient creates a new redis client
func NewClient(host string, port uint16, db uint8, password string) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     strings.Join([]string{host, fmt.Sprint(port)}, ":"),
		Password: password,
		DB:       int(db),
	}

	client := redis.NewClient(options)
	pingStatus := client.Ping()
	err := pingStatus.Err()
	if err != nil {
		return nil, err
	}

	return redis.NewClient(options), nil
}
