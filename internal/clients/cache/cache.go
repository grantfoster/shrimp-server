package cache

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

// GetInstance instantiates or simply returns the existing Redis client.
// Good explanation of the sync.Once usage here at http://marcio.io/2015/07/singleton-pattern-in-go/
func GetInstance() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			Protocol: 2,
		})
	})
	return client
}
