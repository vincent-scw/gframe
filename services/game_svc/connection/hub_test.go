package connection

import (
	"strconv"
	"testing"
	"time"
)

func TestRegisterToHub(t *testing.T) {
	hub := NewHub()
	count := 10
	for i := 0; i < count; i++ {
		hub.register <- &Client{ID: strconv.Itoa(i), send: make(chan *Message)}
	}
	time.Sleep(time.Millisecond * time.Duration(100))

	if len(hub.clients) != count {
		t.Errorf("%d clients should be registered, but was %d.", count, len(hub.clients))
	}

	for j := count - 1; j >= 0; j-- {
		hub.unregister <- strconv.Itoa(j)
	}
	time.Sleep(time.Millisecond * time.Duration(100))
	if len(hub.clients) != 0 {
		t.Error("All clients should be unregistered.")
	}
}
