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

func (c *Client) start(queueRequests chan *Request, numAccounts int) {
	for CONTINUE {
		delay := rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
		r := &Request{}
		c.raffleAccounts(r, numAccounts)
		c.raffleOperation(r)
		queueRequests <- r
	}
}

func (c *Client) raffleAccounts(r *Request, numAccounts int) {
	acc1 := rand.Intn(numAccounts)
	acc2 := rand.Intn(numAccounts)
	r.account1 = acc1
	r.account2 = acc2
}

func (c *Client) raffleOperation(r *Request) {
	numOperations := 3
	r.operation = rand.Intn(numOperations) + 1
	value := 1000
	r.amount = rand.Intn(value*2) - value // Entre -1000 a 1000
}
