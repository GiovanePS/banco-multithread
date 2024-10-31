package server

import (
	"fmt"

	"github.com/GiovanePS/banco-multithread/bank"
)

const (
	DepositarOuSacar int = iota + 1
	Transferir
	BalancoGeral
)

type Worker struct{}

func newWorker(id int) *Worker {
	return &Worker{}
}

func (w *Worker) runJob(bank *bank.Bank, request *Request) {
	switch request.operation {
	case DepositarOuSacar:
		fmt.Println("Depositando ou sancando...")
		// bank.DepositarOuSacar()

	case Transferir:
		fmt.Println("Transferindo...")
		// bank.Transferir()

	case BalancoGeral:
		fmt.Println("Balançeando geral...")
		// bank.BalancoGeral()
	}
}
