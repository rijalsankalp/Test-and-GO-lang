package pointer

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New(" not sufficient fund for withdrawal")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func(w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func(w *Wallet) Balance() Bitcoin {
	return w.balance
}

func(w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	} else {
		w.balance -= amount
		return nil
	}
}

type Stringer interface {
	String() string
}

func(w Bitcoin) String() string {
	return fmt.Sprintf("%d BitCoins", w)
}