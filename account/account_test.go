package account

import "testing"

func TestSaldo(t *testing.T) {
	newAccount := NewAccount(1)
	got := newAccount.Saldo()
	want := 0.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDeposito(t *testing.T) {}

func TestSaque(t *testing.T) {}
