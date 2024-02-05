package test

import (
	"fmt"
	"main.go/tuuz/Redis"
	"testing"
)

func BenchmarkName(b *testing.B) {
	var ps Redis.PubSub
	go func() {
		for message := range ps.Subscribe("test") {
			fmt.Println(message.Payload)
		}
	}()

	ps.Publish("test", []byte("test"))
}
