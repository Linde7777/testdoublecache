package main

import (
	"fmt"
	"testing"
)

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
