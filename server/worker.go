package server

import (
	"fmt"
	"sync"

	"github.com/GiovanePS/banco-multithread/bank"
)

const (
	DepositarOuSacar int = iota + 1
	Transferir
	BalancoGeral
)

type Worker struct {
	mutex sync.Mutex
}

func newWorker() *Worker {
	return &Worker{}
}

func (w *Worker) runJob(bank *bank.Bank, request Request) {
	switch request.operation {
	case DepositarOuSacar:
		fmt.Println("Depositando ou sacando...")
		// bank.DepositarOuSacar(request.account1, float64(request.amount))

	case Transferir:
		fmt.Println("Transferindo...")
		// bank.Transferir(request.account1, request.account2, float64(request.amount))

	case BalancoGeral:
		fmt.Println("Balançeando geral...")
		// bank.BalancoGeral()

	default:
		fmt.Println("Outra operação")
	}
}
