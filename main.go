package main

import (
	"fmt"

	"github.com/GiovanePS/banco-multithread/bank"
)

func main() {
	bank := bank.NewBank()
	bank.CreateAccount()
	err := bank.DepositarOuSacar(1, 4000.50)
	if err != nil {
		fmt.Println(err)
	}

	bank.CreateAccount()
	bank.Transferir(1, 2, 2500)

	err = bank.BalancoGeral()
	if err != nil {
		fmt.Println(err)
	}
}
