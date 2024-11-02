package server

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Test enqueue", func(t *testing.T) {
		queue := &QueueRequests{}
		for i := range 2 {
			req := &Request{account1: i}
			queue.Enqueue(*req)
		}
		fmt.Println(queue)
	})
}
