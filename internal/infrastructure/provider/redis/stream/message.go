package stream

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

func XMessageToString(m *redis.XMessage) string {
	if m == nil {
		return ""
	}
	str := "[" + m.ID + "]\n"
	for key, value := range m.Values {
		str = str + key + " => " + fmt.Sprint(value) + "\n"
	}
	str = "\n"

	return str
}
