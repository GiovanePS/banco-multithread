package bank

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	bank := NewBank()
	want := bank.CreateAccount().Id()
	account, err := bank.GetAccount(0)
	if err != nil {
		t.Fatal(err)
	}

	got := account.Id()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDepositarOuSacar(t *testing.T) {
	bank := NewBank()
	acc := bank.CreateAccount()

	t.Run("Sucesso ao depositar", func(t *testing.T) {
		err := bank.DepositarOuSacar(1, 10)
		if err != nil {
			t.Fatal(err)
		}

		got := acc.Saldo()
		want := 10.0

		if got != want {
			t.Errorf("got '%f' want '%f'", got, want)
		}
	})

	t.Run("Falha ao depositar em uma conta que não existe", func(t *testing.T) {
		err := bank.DepositarOuSacar(2, 0)
		if err == nil {
			t.Errorf("Erro não retornado em uma operação inválida.")
		}
	})

	t.Run("Sucesso ao sacar", func(t *testing.T) {
		err := bank.DepositarOuSacar(1, -5)
		if err != nil {
			t.Fatal(err)
		}

		got := acc.Saldo()
		want := 5.0

		if got != want {
			t.Errorf("got '%f' want '%f'", got, want)
		}
	})

	t.Run("Falha ao sacar um valor maior do que o disponível na conta", func(t *testing.T) {
		err := bank.DepositarOuSacar(1, -1000)
		if err == nil {
			t.Errorf("Erro não retornado ao sacar um valor maior do que o disponível na conta.")
		}
	})
}
