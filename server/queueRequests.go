package server

import (
	"fmt"
	"sync"
)

type QueueRequests struct {
	Head   *NodeRequests
	Tail   *NodeRequests
	mutex  sync.Mutex
	Length int
}

type NodeRequests struct {
	Request *Request
	Next    *NodeRequests
	Back    *NodeRequests
}

func (qr *QueueRequests) Enqueue(newRequest Request) {
	qr.mutex.Lock()
	defer qr.mutex.Unlock()

	qr.Length++
	if qr.Head == nil {
		newNode := &NodeRequests{Request: &newRequest}
		qr.Head = newNode
		qr.Tail = newNode
		return
	}

	newNode := &NodeRequests{Request: &newRequest}
	newNode.Next = qr.Head
	qr.Head.Back = newNode
	qr.Head = newNode
}

func (qr *QueueRequests) Dequeue() (Request, error) {
	qr.mutex.Lock()
	defer qr.mutex.Unlock()

	// Se não houver nenhum node
	if qr.Head == nil {
		return Request{}, fmt.Errorf("Não há nodes para remover.")
	}

	qr.Length--
	// Se houver apenas um Node
	if qr.Head == qr.Tail {
		request := *qr.Head.Request
		qr.Head = nil
		qr.Tail = nil
		return request, nil
	}

	temp := qr.Tail
	qr.Tail = qr.Tail.Back
	qr.Tail.Next = nil
	temp.Back = nil
	request := *temp.Request
	return request, nil
}

func (qr *QueueRequests) String() string {
	qr.mutex.Lock()
	defer qr.mutex.Unlock()
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
