package server

import (
	"fmt"
	"sync"
)

type QueueRequests struct {
	Head         *NodeRequests
	Tail         *NodeRequests
	queueMutex   sync.Mutex
	EnqueueMutex sync.Mutex
	DequeueMutex sync.Mutex
	Length       int
}

type NodeRequests struct {
	Request *Request
	Next    *NodeRequests
	Back    *NodeRequests
}

func (qr *QueueRequests) Enqueue(newRequest *Request) {
	// Menor que 2 pois variáveis Head e Tail são mexidas nos dois métodos.
	if qr.Length <= 2 {
		qr.queueMutex.Lock()
		defer qr.queueMutex.Unlock()
	} else {
		qr.EnqueueMutex.Lock()
		defer qr.EnqueueMutex.Unlock()
	}

	if qr.Head == nil {
		newNode := &NodeRequests{Request: newRequest}
		qr.Head = newNode
		qr.Tail = newNode
		return
	}

	newNode := &NodeRequests{Request: newRequest}
	newNode.Next = qr.Head
	qr.Head.Back = newNode
	qr.Head = newNode
	qr.Length++
}

func (qr *QueueRequests) Dequeue() (*Request, error) {
	if qr.Length <= 2 {
		qr.queueMutex.Lock()
		defer qr.queueMutex.Unlock()
	} else {
		qr.DequeueMutex.Lock()
		defer qr.DequeueMutex.Unlock()
	}

	// Se não houver nenhum node
	if qr.Head == nil {
		return nil, fmt.Errorf("Não há nodes para remover.")
	}

	qr.Length--
	// Se houver apenas um Node
	if qr.Head == qr.Tail {
		request := qr.Head.Request
		qr.Head = nil
		qr.Tail = nil
		return request, nil
	}

	temp := qr.Tail
	qr.Tail = qr.Tail.Back
	qr.Tail.Next = nil
	temp.Back = nil
	return temp.Request, nil
}

func (qr *QueueRequests) String() string {
	output := "["
	runner := qr.Head

	// Queue vazia
	if runner == nil {
		output += "]"
		return output
	}

	for {
		output += fmt.Sprintf("%d", runner.Request.account1)
		if runner.Next != nil {
			output += ", "
			runner = runner.Next
		} else {
			break
		}
	}

	output += "]"
	return output
}
