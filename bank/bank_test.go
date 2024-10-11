package bank

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	bank := NewBank()

	t.Run("Testar sucesso ao pegar uma conta", func(t *testing.T) {
		want := bank.CreateAccount().id
		account, err := bank.GetAccount(1)
		if err != nil {
			t.Fatal(err)
		}

		got := account.id

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

		got := acc.saldo
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

		got := acc.saldo
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

func TestTransferir(t *testing.T) {
	t.Run("Sucesso ao transferir", func(t *testing.T) {
		bank := NewBank()
		acc1 := bank.CreateAccount()
		err := bank.DepositarOuSacar(1, 10)
		if err != nil {
			t.Fatal(err)
		}

		acc2 := bank.CreateAccount()
		err = bank.DepositarOuSacar(2, 10)
		if err != nil {
			t.Fatal(err)
		}

		bank.Transferir(1, 2, 5)
		got := acc1.saldo
		want := 5.0

		if got != want {
			t.Errorf("Saldo do remetente da transferência incorreto: got: %f, want: %f", got, want)
		}

		got = acc2.saldo
		want = 15.0

		if got != want {
			t.Errorf("Saldo do destinatário da transferência incorreto: got: %f, want: %f", got, want)
		}
	})
}
