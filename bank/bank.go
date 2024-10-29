package bank

import (
	"fmt"
	"time"
)

// Tempo em milisegundos
var (
	delayOfDepositarOuSacar = 200
	delayOfTransferir       = 300
	delayOfBalancoGeral     = 500
)

// O serviço deve manter informação sobre contas de usuários. Cada conta
// de usuário tem um identificador da conta (um número inteiro positivo)
// e um saldo atual (um número real, que pode ser positivo ou negativo).
type Bank struct {
	id               int
	headlistAccounts *Node
	idCountContas    int
}

type Node struct {
	account *Account
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

func (b *Bank) CreateAccount() *Account {
	if b.headlistAccounts.account == nil {
		b.idCountContas++
		newAccount := newAccount(b.idCountContas)
		b.headlistAccounts.account = newAccount
		return newAccount
	}

	r := b.headlistAccounts
	for {
		if r.next == nil {
			b.idCountContas++
			newAccount := newAccount(b.idCountContas)
			r.next = &Node{account: newAccount}
			return newAccount
		}

		r = r.next
	}
}

func (b *Bank) GetAccount(accountId int) (*Account, error) {
	if b.headlistAccounts.account == nil {
		return nil, fmt.Errorf("Nenhuma conta existente.")
	}

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
	time.Sleep(time.Millisecond * time.Duration(delayOfDepositarOuSacar))
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
	time.Sleep(time.Millisecond * time.Duration(delayOfTransferir))
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

// Esta operação gera um balanço geral de todas as contas, imprimindo na tela cada
// conta e o seu respectivo valor no momento em que a operação foi inicializada.
// Note que o balanço geral apresenta uma “fotografia” instantânea do estado das contas.
func (b *Bank) BalancoGeral() error {
	time.Sleep(time.Millisecond * time.Duration(delayOfBalancoGeral))
	if b.headlistAccounts.account == nil {
		return fmt.Errorf("Nenhuma conta existente.")
	}

	r := b.headlistAccounts
	for {
		fmt.Printf("Account ID: %d\tAccount Balance: %.2f\n", r.account.id, r.account.saldo)

		if r.next == nil {
			return nil
		}

		r = r.next
	}
}
