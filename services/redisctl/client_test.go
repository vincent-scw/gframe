package redisctl

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestPubSub(t *testing.T) {
// 	cli := NewRedisClient("localhost:6379")

// 	channel := "testchan"
// 	go cli.Subscribe(channel, func(content string) string {
// 		fmt.Println(content)
// 		return content
// 	})

// 	for i := 0; i < 10; i++ {
// 		cli.Publish(channel, fmt.Sprintf("msg %d", i))
// 	}

// 	time.AfterFunc(time.Second, func() { cli.Close() })
// }
