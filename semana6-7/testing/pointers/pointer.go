package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
