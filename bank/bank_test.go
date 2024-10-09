package bank

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	bank := NewBank()

	t.Run("Testar sucesso ao pegar uma conta", func(t *testing.T) {
		want := bank.CreateAccount().Id()
		account, err := bank.GetAccount(1)
		if err != nil {
			t.Fatal(err)
		}

		got := account.Id()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Testar falha ao tentar pegar uma conta inexistente", func(t *testing.T) {
		_, err := bank.GetAccount(1000)
		if err == nil {
			t.Errorf("Erro não retornado ao pegar um usuário inexistente.")
		}
	})
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
		err := bank.DepositarOuSacar(1000, 0)
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
