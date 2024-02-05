package main

import (
	"fmt"
	"main.go/tuuz/Redis"
	"testing"
)

func BenchmarkName(b *testing.B) {
	var ps Redis.Pubsub[byte]
	go func() {
		for message := range ps.Subscribe("test") {
			fmt.Println(message)
		}
	}()

	ps.Publish("test", []byte("test"))
}
