package server

import (
	"math/rand"
	"time"
)

type Request struct {
	account1  int
	account2  int // Se operação necessitar de outra conta
	operation int
	amount    float64
}

type Client struct{}

func (c *Client) start(chRequests chan Request) {
	for {
		delay := rand.Intn(3) + 1
		time.Sleep(time.Duration(delay))
		r := Request{}
		c.raffAccounts(r)
		chRequests <- r
		time.Sleep(time.Duration(delay))
	}
}

func (c *Client) raffAccounts(numAccounts int, r Request) {
	acc1 := rand.Intn(numAccounts)
	acc2 := rand.Intn(numAccounts)
	r.account1 = acc1
	r.account2 = acc2
}

func (c *Client) raffOperation() int {
	numOperations := 3
	return rand.Intn(numOperations)
}
