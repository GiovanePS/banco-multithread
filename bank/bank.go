package bank

import (
	"fmt"
)

type Bank struct {
	id               int
	headlistAccounts *Node
}

type Node struct {
	account *Account
	next    *Node
}

type Account struct {
	id    int
	saldo float64
}

func newAccount(idContador int) *Account {
	newAccount := &Account{idContador, 0.0}
	idContador++
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

func NewBank() *Bank {
	newBank := &Bank{
		headlistAccounts: &Node{},
	}

	return newBank
}

func (b *Bank) Id() int {
	return b.id
}

var idCountContas = 1

func (b *Bank) CreateAccount() *Account {
	if b.headlistAccounts.account == nil {
		newAccount := newAccount(idCountContas)
		b.headlistAccounts.account = newAccount
		return newAccount
	}

	r := b.headlistAccounts
	for {
		if r.next == nil {
			newAccount := newAccount(idCountContas)
			r.next = &Node{account: newAccount}
			return newAccount
		}

		r = r.next
	}
}

func (b *Bank) GetAccount(accountId int) (*Account, error) {
	r := b.headlistAccounts
	for {
		if r.account.id == accountId {
			return r.account, nil
		}

		if r.next == nil {
			return nil, fmt.Errorf("Conta com identificador %d inexistente", accountId)
		}

		r = r.next
	}
}

// Esta operação recebe um identificador de conta (um número inteiro positivo) e
// o valor de depósito (um número real, que pode ser positivo ou negativo). Note
// que esta operação pode ser executada tanto para depósitos quanto saques,
// dependendo se o valor de depósito é positivo ou negativo;
func (b *Bank) DepositarOuSacar(accountId int, value float64) error {
	acc, err := b.GetAccount(accountId)
	if err != nil {
		return err
	}

	if value <= 0 {
		err := acc.sacar(value)
		if err != nil {
			return err
		}

		return nil
	}

	err = acc.depositar(value)
	if err != nil {
		return err
	}

	return nil
}

// Dadas duas contas bancárias, origem e destino, e um valor de transferência, esta
// operação deve debitar o valor de transferência da conta de origem e somar este
// valor na conta destino;
func (b *Bank) Transferir(sourceAccount int, destAccount int, value float64) error {
	accSource, err := b.GetAccount(sourceAccount)
	if err != nil {
		return err
	}

	accDest, err := b.GetAccount(destAccount)
	if err != nil {
		return err
	}

	err = accSource.sacar(value)
	if err != nil {
		return err
	}

	err = accDest.depositar(value)
	if err != nil {
		accSource.depositar(value)
		return err
	}

	return nil
}
