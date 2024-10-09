package main

import (
	"fmt"

	"github.com/GiovanePS/banco-multithread/bank"
)

func main() {
	banco := bank.NewBank()
	myAccount := banco.CreateAccount()
	fmt.Printf("%+v\n", myAccount.Saldo())
}
