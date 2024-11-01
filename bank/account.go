package bank

import (
	"fmt"
	"sync"
)

type Account struct {
	id    int
	saldo float64
	mutex sync.Mutex
}

func newAccount(idContador int) *Account {
	newAccount := &Account{id: idContador, saldo: 0.0}
	return newAccount
}

func (a *Account) depositar(valor float64) error {
	if valor < 0 {
		return fmt.Errorf("Só é possível realizar depósitos de valores positivos")
	}

	a.saldo += valor
	return nil
}

func (a *Account) sacar(valor float64) error {
	// absolute value
	if valor < 0 {
		valor = -valor
	}

	if a.saldo < valor {
		return fmt.Errorf("Saldo insuficiente para sacar R$ %v", valor)
	}

	a.saldo -= valor
	return nil
}
