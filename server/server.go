package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

const maxRequests = 50

var CONTINUE = true

type Server struct {
	bank          *bank.Bank
	workers       []*Worker
	numWorkers    int
	clients       []*Client
	queueRequests *QueueRequests
	numAccounts   int
}

func CreateServerThread(numAgents int) *Server {
	s := &Server{
		bank:          bank.NewBank(),
		workers:       make([]*Worker, numAgents),
		numWorkers:    numAgents,
		clients:       make([]*Client, numAgents),
		queueRequests: &QueueRequests{},
		numAccounts:   10,
	}

	s.createThreadPool(numAgents)
	s.createBankAccounts(numAgents)
	return s
}

func (s *Server) createBankAccounts(numAccounts int) {
	for range numAccounts {
		s.bank.CreateAccount()
	}
}

func (s *Server) createThreadPool(numWorkers int) {
	for i := range numWorkers {
		s.workers[i] = newWorker()
	}
}

func (s *Server) StartServerThread() {
	var wg sync.WaitGroup

	// Setting up clients
	for _, client := range s.clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client.start(s.queueRequests, s.numAccounts)
		}()
	}

	requestCount := 0

	// Setting up workers
	turnCount := 0
	for {
		request, err := s.queueRequests.Dequeue()
		if err == nil {
			worker := s.workers[turnCount%s.numWorkers]
			for range 10 {
				if worker.Mutex.TryLock() {
					wg.Add(1)
					go func() {
						defer wg.Done()
						worker.runJob(s.bank, request)
						worker.Mutex.Unlock()
					}()
					turnCount++
					break
				}
			}

			requestCount++
		}

		if requestCount == maxRequests {
			CONTINUE = false
			break
		}
	}

	wg.Wait()
	CONTINUE = true
}
