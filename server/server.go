package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

var CONTINUE = true

type Server struct {
	bank          *bank.Bank
	workers       chan *Worker
	clients       []*Client
	queueRequests chan *Request
	numAccounts   int
}

func CreateServerThread(numAgents int) *Server {
	s := &Server{
		bank:          bank.NewBank(),
		workers:       make(chan *Worker, numAgents),
		clients:       make([]*Client, numAgents),
		queueRequests: make(chan *Request, numAgents),
		numAccounts:   10,
	}

	s.createThreadPool(numAgents)
	return s
}

func (s *Server) createBankAccounts(numAccounts int) {
	for range numAccounts {
		s.bank.CreateAccount()
	}
}

func (s *Server) createThreadPool(numWorkers int) {
	for i := range numWorkers {
		s.workers <- newWorker(i + 1)
	}
}

func (s *Server) StartServerThread() {
	var wg sync.WaitGroup

	for _, client := range s.clients {
		wg.Add(1)
		go func(client *Client) {
			defer wg.Done()
			client.start(s.queueRequests, s.numAccounts)
		}(client)
	}
	maxRequests := 25
	requestCount := 0

	// Receive requests
	for {
		request := <-s.queueRequests

		wg.Add(1)
		go func() {
			defer wg.Done()
			worker := <-s.workers
			worker.runJob(s.bank, request)
			s.workers <- worker
		}()

		requestCount++
		if requestCount == maxRequests {
			CONTINUE = false
			break
		}
	}

	wg.Wait()
}
