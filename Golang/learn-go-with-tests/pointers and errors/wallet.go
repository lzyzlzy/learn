package wallet

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= Bitcoin(amount)
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
