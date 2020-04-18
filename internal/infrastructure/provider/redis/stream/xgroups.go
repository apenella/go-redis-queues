package stream

import "github.com/go-redis/redis/v7"

func existsXGroup(client *redis.Client, group string) bool {
	xinfogroupscmd := client.XInfoGroups(group)
	
	if xinfogroupscmd == nil {
		return false
	}

	return true
}