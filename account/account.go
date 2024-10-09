package account

import (
	"fmt"
)

type Account struct {
	id    int
	saldo float64
}

var idContador = 0

func NewAccount() *Account {
	newAccount := &Account{idContador, 0.0}
	idContador++
	return newAccount
}

func (a *Account) Depositar(valor float64) error {
	if valor < 0 {
		return fmt.Errorf("Só é possível realizar depósitos de valores positivos")
	}

	a.saldo += valor
	return nil
}

func (a *Account) Sacar(valor float64) error {
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

func (a *Account) Id() int {
	return a.id
}

func (a *Account) Saldo() float64 {
	return a.saldo
}
