package bank

import (
	"fmt"

	"github.com/GiovanePS/banco-multithread/account"
)

type Bank struct {
	id               int
	headlistAccounts *Node
}

type Node struct {
	account *account.Account
	next    *Node
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

func (b *Bank) CreateAccount() *account.Account {
	if b.headlistAccounts.account == nil {
		newAccount := account.NewAccount()
		b.headlistAccounts.account = newAccount
		return newAccount
	}

	r := b.headlistAccounts
	for {
		if r.next == nil {
			newAccount := account.NewAccount()
			r.next = &Node{account: newAccount}
			return newAccount
		}

		r = r.next
	}
}

func (b *Bank) GetAccount(accountId int) (*account.Account, error) {
	r := b.headlistAccounts
	for {
		if r.account.Id() == accountId {
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
func (b *Bank) DepositarOuSacar(accountId int, valor float64) error {
	acc, err := b.GetAccount(accountId)
	if err != nil {
		return err
	}

	if valor <= 0 {
		err := acc.Sacar(valor)
		if err != nil {
			return err
		}

		return nil
	}

	err = acc.Depositar(valor)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bank) Transferir(sourceAccount *account.Account, destAccount *account.Account) error {
	return nil
}
