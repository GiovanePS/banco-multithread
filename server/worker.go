package server

import (
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

// Job representa uma operação de banco.

type Worker struct {
	id   int
	cond *sync.Cond
}

func newWorker(id int) *Worker {
	newWorker := &Worker{
		id:   id,
		cond: sync.NewCond(&sync.Mutex{}),
	}

	return newWorker
}

func (w *Worker) start(bank *bank.Bank, jobQueue <-chan Request) {
	for job := range jobQueue {
		w.cond.L.Lock()
		w.cond.Signal()
		w.cond.L.Unlock()
	}
}

func (w *Worker) runJob(bank bank.Bank, operation int) {
	switch operation {
	case 1:
		bank.DepositarOuSacar()

	case 2:
		bank.Transferir()

	case 3:
		bank.BalancoGeral()
	}
}
