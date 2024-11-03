package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

type Server struct {
	bank          *bank.Bank
	workers       []*Worker
	numWorkers    int
	clients       []*Client
	numClients    int
	queueRequests *QueueRequests
	numAccounts   int
	numRequests   int
	serviceTime   int
}

func CreateServerThread(numWorkers, numClients, numRequests, serviceTime int) *Server {
	s := &Server{
		bank:          bank.NewBank(serviceTime),
		workers:       make([]*Worker, numWorkers),
		numWorkers:    numWorkers,
		clients:       make([]*Client, numClients),
		numClients:    numClients,
		queueRequests: &QueueRequests{},
		numAccounts:   numClients,
		numRequests:   numRequests,
		serviceTime:   serviceTime,
	}

	s.createThreadPool(numWorkers)
	s.createBankAccounts(numClients)
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
	getRequest := func() (Request, error) {
		request, err := s.queueRequests.Dequeue()
		return request, err
	}

	turnWorker := 0
	getAvailableWorker := func() *Worker {
		for {
			worker := s.workers[turnWorker%s.numWorkers]
			turnWorker++
			if worker.mutex.TryLock() {
				return worker
			}
		}
	}

	// Starting clients
	var wg sync.WaitGroup
	for _, client := range s.clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < s.numRequests; i++ {
				client.send(s.queueRequests, s.numAccounts)
			}
		}()
	}

	// Stating handle
	wg.Add(1)
	go func() {
		defer wg.Done()
		for handledRequests := 0; handledRequests < s.numClients*s.numRequests; handledRequests++ {
			var request Request
			for {
				var err error
				request, err = getRequest()
				// Há requests para executar
				if err == nil {
					break
				}
			}

			worker := getAvailableWorker()
			wg.Add(1)
			go func() {
				defer wg.Done()
				worker.runJob(s.bank, request)
				worker.mutex.Unlock()
			}()
		}
	}()

	wg.Wait()
}
