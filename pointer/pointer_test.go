package pointer

import (
	"testing"
)


func TestPointer(t *testing.T) {
	t.Run("testing Deposit", func(t *testing.T) {
		wallet :=Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertbalance(t, wallet, Bitcoin(10))
	})

	t.Run("testing Withdraw", func(t *testing.T) {
		wallet :=Wallet{balance:Bitcoin(30)}
		error_return := wallet.Withdraw(Bitcoin(29))

		assertnoerror(t, error_return)
		assertbalance(t, wallet, Bitcoin(1))
	})

	t.Run("testing Withdraw when less", func(t *testing.T) {
		startingbalance := Bitcoin(10)
		wallet := Wallet{balance:startingbalance}
		error_return :=wallet.Withdraw(Bitcoin(40))

		asserterror(t, error_return, ErrorInsufficientFunds)
		assertbalance(t, wallet, Bitcoin(10))
		
	})
}

func assertbalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("want %s, but got %s", want, got)
	}
}

func asserterror(t testing.TB, err, want_error error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't got one,")
	}
	if err.Error() != want_error.Error() {
		t.Errorf("got %s, want %s", err, want_error)
	}
}

func assertnoerror(t testing.TB, err error) {
	t.Helper()
	if err !=nil {
		t.Error("didn't wanted an error, but got one.")
	}
}