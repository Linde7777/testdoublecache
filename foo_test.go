package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"
	"testing"
)

func BenchmarkRedis(b *testing.B) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
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
	mcClient = memcache.New("localhost:11211")

	for i := 0; i < b.N; i++ {
		for j := 0; j < keyCount; j++ {
			_, err := mcClient.Get(fmt.Sprintf("key%d", j))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
