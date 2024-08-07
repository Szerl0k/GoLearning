package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount // == (*w).balance	(Go does automatic dereference)
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("address of balance in Deposit is %p \n\n", &w.balance)
	return w.balance // == (*w).balance	(Go does automatic dereference)
}

// ErrInsufficientFunds global (to the package) variable
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil

}
