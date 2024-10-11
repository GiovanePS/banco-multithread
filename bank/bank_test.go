package bank

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	bank := NewBank()

	t.Run("Sucesso ao pegar uma conta", func(t *testing.T) {
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

	t.Run("Falhar ao tentar pegar uma conta inexistente", func(t *testing.T) {
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

	t.Run("Falhar ao depositar em uma conta que não existe", func(t *testing.T) {
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

	t.Run("Falhar ao sacar um valor maior do que o disponível na conta", func(t *testing.T) {
		err := bank.DepositarOuSacar(1, -1000)
		if err == nil {
			t.Errorf("Erro não retornado ao sacar um valor maior do que o disponível na conta.")
		}
	})
}

func TestTransferir(t *testing.T) {
	setup := func() (bank *Bank, acc1 *Account, acc2 *Account) {
		bank = NewBank()
		acc1 = bank.CreateAccount()
		err := bank.DepositarOuSacar(1, 10)
		if err != nil {
			t.Fatal(err)
		}

		acc2 = bank.CreateAccount()
		err = bank.DepositarOuSacar(2, 10)
		if err != nil {
			t.Fatal(err)
		}

		return bank, acc1, acc2
	}

	checkErrors := func(t *testing.T, sourceGot float64, sourceWant float64, destGot float64, destWant float64) {
		if sourceGot != sourceWant {
			t.Errorf("Saldo do remetente da transferência incorreto: got: %f, want: %f", sourceGot, sourceWant)
		}

		if destGot != destWant {
			t.Errorf("Saldo do destinatário da transferência incorreto: got: %f, want: %f", destGot, destWant)
		}
	}

	t.Run("Sucesso ao transferir", func(t *testing.T) {
		bank, acc1, acc2 := setup()
		err := bank.Transferir(1, 2, 5)
		if err != nil {
			t.Fatal(err)
		}

		sourceGot := acc1.saldo
		sourceWant := 5.0
		destGot := acc2.saldo
		destWant := 15.0

		checkErrors(t, sourceGot, sourceWant, destGot, destWant)
	})

	t.Run("Falhar ao transferir com conta destino sem saldo suficiente", func(t *testing.T) {
		bank, _, _ := setup()
		err := bank.Transferir(1, 2, 15) // Saldo insuficiente para source
		if err == nil {
			t.Errorf("Erro não retornado ao sacar um valor maior do que o disponível na conta.")
		}
	})
}
