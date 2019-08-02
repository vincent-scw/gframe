package redisctl

import (
	"fmt"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	cli := NewPubSubClient("40.83.112.48:6379")

	channel := "testchan"
	go cli.Subscribe(channel, func(content string) {
		fmt.Println(content)
	})

	for i := 0; i < 10; i++ {
		cli.Publish(channel, fmt.Sprintf("msg %d", i))
	}

	time.AfterFunc(time.Second, func() { cli.Close() })
}
