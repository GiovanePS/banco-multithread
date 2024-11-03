package bank

import (
	"fmt"
	"sync"
	"time"
)

var balancing = false
var cond = sync.NewCond(&sync.Mutex{})

// O serviço deve manter informação sobre contas de usuários. Cada conta
// de usuário tem um identificador da conta (um número inteiro positivo)
// e um saldo atual (um número real, que pode ser positivo ou negativo).
type Bank struct {
	id               int
	headlistAccounts *NodeAccount
	idCountContas    int
	serviceTime      time.Duration // Em segundos
	mutex            sync.Mutex
}

type NodeAccount struct {
	account *Account
	next    *NodeAccount
}

func NewBank(serviceTime int) *Bank {
	newBank := &Bank{
		headlistAccounts: &NodeAccount{},
		serviceTime:      time.Duration(serviceTime),
	}

	return newBank
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
			r.next = &NodeAccount{account: newAccount}
			return newAccount
		}

		r = r.next
	}
}

func (b *Bank) GetAccount(accountId int) (*Account, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
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
func (b *Bank) DepositarOuSacar(accountId int, valor float64) error {
	time.Sleep(b.serviceTime * time.Millisecond)
	acc, err := b.GetAccount(accountId)
	cond.L.Lock()
	for balancing {
		cond.Wait()
	}
	cond.L.Unlock()
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	if err != nil {
		fmt.Errorf("Erro ao buscar conta: " + err.Error())
		return err
	}

	if valor <= 0 {
		err := acc.sacar(valor)
		if err != nil {
			fmt.Errorf("Erro ao sacar: " + err.Error())
			return err
		}

		valor = -valor
		fmt.Printf("Sacando na conta %d: %f\n", accountId, valor)
		return nil
	}

	err = acc.depositar(valor)
	if err != nil {
		fmt.Errorf("Erro ao depositar: " + err.Error())
		return err
	}

	fmt.Printf("Depositando na conta %d: %f\n", accountId, valor)
	return nil
}

// Dadas duas contas bancárias, origem e destino, e um valor de transferência, esta
// operação deve debitar o valor de transferência da conta de origem e somar este
// valor na conta destino;
func (b *Bank) Transferir(sourceAccount int, destAccount int, valor float64) error {
	time.Sleep(b.serviceTime * time.Millisecond)
	accSource, err := b.GetAccount(sourceAccount)
	if err != nil {
		return err
	}

	accDest, err := b.GetAccount(destAccount)
	if err != nil {
		return err
	}

	cond.L.Lock()
	for balancing {
		cond.Wait()
	}
	cond.L.Unlock()

	// Evitar deadlocks de operações usando contas inversas
	if accSource.id < accDest.id {
		accSource.mutex.Lock()
		defer accSource.mutex.Unlock()
		accDest.mutex.Lock()
		defer accDest.mutex.Unlock()
	} else {
		accDest.mutex.Lock()
		defer accDest.mutex.Unlock()
		accSource.mutex.Lock()
		defer accSource.mutex.Unlock()
	}

	err = accSource.sacar(valor)
	if err != nil {
		return fmt.Errorf("Erro ao transferir: %s", err)
	}

	err = accDest.depositar(valor)
	if err != nil {
		accSource.depositar(valor)
		return fmt.Errorf("Erro ao transferir: %s", err)
	}

	if valor < 0 {
		valor = -valor
	}
	fmt.Printf("Transferência: %d -> %d : %f\n", accSource.id, accDest.id, valor)
	return nil
}

// Esta operação gera um balanço geral de todas as contas, imprimindo na tela cada
// conta e o seu respectivo valor no momento em que a operação foi inicializada.
// Note que o balanço geral apresenta uma “fotografia” instantânea do estado das contas.
func (b *Bank) BalancoGeral() error {
	time.Sleep(b.serviceTime * time.Millisecond)
	if b.headlistAccounts.account == nil {
		return fmt.Errorf("Nenhuma conta existente.")
	}

	balancing = true
	r := b.headlistAccounts
	fmt.Println("Balanço geral: ")
	for {
		r.account.mutex.Lock()
		fmt.Printf("Account ID: %d\tAccount Balance: %.2f\n", r.account.id, r.account.saldo)
		r.account.mutex.Unlock()

		if r.next == nil {
			balancing = false
			cond.Broadcast()
			return nil
		}

		r = r.next
	}
}
