package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"
	"testing"
)

var redisClient *redis.Client
var mcClient *memcache.Client

const keyCount = 1000

func setup() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	mcClient = memcache.New("localhost:11211")

	// Populate Redis and Memcached with 100 key-value pairs
	for i := 0; i < keyCount; i++ {
		err := redisClient.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), 0).Err()
		if err != nil {
			panic(err)
		}

		err = mcClient.Set(&memcache.Item{Key: fmt.Sprintf("key%d", i), Value: []byte(fmt.Sprintf("value%d", i))})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < keyCount; j++ {
			_, err := redisClient.Get(fmt.Sprintf("key%d", j)).Result()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

func BenchmarkMemcached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < keyCount; j++ {
			_, err := mcClient.Get(fmt.Sprintf("key%d", j))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

func main() {
	setup()
}
