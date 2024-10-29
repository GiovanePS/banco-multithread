package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

type Server struct {
	bank       *bank.Bank
	workers    []*Worker
	clients    []*Client
	chRequests chan Request
}

func CreateServerThread(numAgents int) *Server {
	s := &Server{
		bank:       bank.NewBank(),
		workers:    make([]*Worker, numAgents),
		clients:    make([]*Client, numAgents),
		chRequests: make(chan Request, numAgents),
	}

	s.createThreadPool(numAgents)
	return s
}

func (s *Server) createAccounts(numAccounts int) {
	for range numAccounts {
		s.bank.CreateAccount()
	}
}

func (s *Server) createThreadPool(numWorkers int) {
	for i := range numWorkers {
		s.workers[i] = newWorker(i + 1)
	}
}

func (s *Server) InitServerThread() {
	var wait sync.WaitGroup

	for _, worker := range s.workers {
		wait.Add(1)
		go worker.start(s.bank, s.chRequests)
	}

	for _, client := range s.clients {
		wait.Add(1)
		go client.start(s.chRequests)
	}

	// Receive requests
	for {
		if true {
			break
		}
	}

	go func() {
		wait.Wait()
	}()
}

func (s *Server) Start(wait *sync.WaitGroup) {
	defer wait.Done()
}

// Start inicia o worker para processar Jobs
// func (w *Worker) Start(results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case job := <-w.jobChan: // Recebe um Job
// 			sum := 0
// 			for i := job.start; i <= job.end; i++ {
// 				sum += i
// 			}
// 			fmt.Printf("Worker %d processando soma de %d a %d = %d\n", w.id, job.start, job.end, sum)
// 			results <- sum
// 		case <-w.done: // Finaliza o worker
// 			return
// 		}
// 	}
// }
