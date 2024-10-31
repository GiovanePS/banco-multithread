package server

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("Testing requests from client", func(t *testing.T) {
		channel := make(chan *Request)
		client := Client{}
		for range 10 {
			go client.start(channel, 10)
			fmt.Println(<-channel)
		}
	})
}

func TestServer(t *testing.T) {
	server := Server{}
	server.createThreadPool(10)
}
