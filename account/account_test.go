package account

import "testing"

func TestSaldo(t *testing.T) {
	newAccount := NewAccount()
	got := newAccount.Saldo()
	want := 0.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDeposito(t *testing.T) {}

func TestSaque(t *testing.T) {}
