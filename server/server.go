package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

var cond = sync.NewCond(&sync.Mutex{})

type Server struct {
	bank           *bank.Bank
	workers        []*Worker
	numWorkers     int
	clients        []*Client
	clientsRunning int
	numClients     int
	queueRequests  *QueueRequests
	numAccounts    int
	numRequests    int
	serviceTime    int
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

	queueIsEmpty := func() bool {
		s.queueRequests.mutex.Lock()
		defer s.queueRequests.mutex.Unlock()
		numReqs := s.queueRequests.Length
		if numReqs > 0 {
			return false
		}

		return true
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
	mutexClientsRunning := sync.Mutex{}
	var wg sync.WaitGroup
	for _, client := range s.clients {
		wg.Add(1)
		mutexClientsRunning.Lock()
		s.clientsRunning++
		mutexClientsRunning.Unlock()
		go func() {
			defer wg.Done()
			for i := 0; i < s.numRequests; i++ {
				client.send(s.queueRequests, s.numAccounts)
			}
			mutexClientsRunning.Lock()
			s.clientsRunning--
			mutexClientsRunning.Unlock()
			// Notificar que finalizou
			cond.Signal()
		}()
	}

	sendBalanceRequest := func() {
		r := Request{operation: 3}
		s.queueRequests.Enqueue(r)
	}

	clientAreNotRunning := func() bool {
		mutexClientsRunning.Lock()
		defer mutexClientsRunning.Unlock()
		if s.clientsRunning > 0 {
			return false
		}

		return true
	}

	// Stating main thread
	countRequests := 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if queueIsEmpty() && clientAreNotRunning() {
				break
			}

			cond.L.Lock()
			for queueIsEmpty() && !clientAreNotRunning() {
				cond.Wait()
			}
			cond.L.Unlock()

			if countRequests%10 == 0 {
				sendBalanceRequest()
			}
			countRequests++

			request, _ := getRequest()
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
