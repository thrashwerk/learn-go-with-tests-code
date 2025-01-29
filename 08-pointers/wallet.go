package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Shitcoin int

func (s Shitcoin) String() string {
	return fmt.Sprintf("%d STC", s)
}

type Wallet struct {
	balance Shitcoin
}

func (w *Wallet) Deposit(n Shitcoin) {
	w.balance += n
}

func (w *Wallet) Balance() Shitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(n Shitcoin) error {
	if n > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= n
	return nil
}
