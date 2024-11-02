package server

import (
	"math/rand"
	"time"
)

type Request struct {
	account1  int
	account2  int // Se operação necessitar de outra conta
	operation int
	amount    int
}

type Client struct{}

func (c *Client) send(qr *QueueRequests, numAccounts int) {
	delay := rand.Intn(5)
	time.Sleep(time.Duration(delay) * time.Second)
	r := &Request{}
	c.raffleAccounts(r, numAccounts)
	c.raffleOperation(r)
	qr.Enqueue(*r)
}

func (c *Client) raffleAccounts(r *Request, numAccounts int) {
	acc1 := rand.Intn(numAccounts) + 1
	acc2 := rand.Intn(numAccounts) + 1

	// Se acc2 for igual à acc1
	for acc2 == acc1 {
		acc2 = rand.Intn(numAccounts) + 1
	}

	r.account1 = acc1
	r.account2 = acc2
}

func (c *Client) raffleOperation(r *Request) {
	numOperations := 3
	r.operation = rand.Intn(numOperations) + 1
	value := 1000
	r.amount = rand.Intn(value*2) - value // Entre -1000 a 1000
}
