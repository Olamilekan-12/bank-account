package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAmount = errors.New("amount must be greater than zero")
var ErrInsufficientFunds = errors.New("insufficient funds")

type Account struct {
	Owner   string
	Balance int
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	return nil

}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.Balance {
		return fmt.Errorf("withdraw %d from balance %d: %v", amount, a.Balance, ErrInsufficientFunds)
	}
	a.Balance -= amount
	return nil
}

func main() {
	myFirstAccount := &Account{}
	err := myFirstAccount.Deposit(100)
	if err != nil {
		fmt.Println(err)
	}
	err = myFirstAccount.Withdraw(500)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Printf("[insufficient funds branch] %v\n", err)
	} else if err != nil {
		fmt.Println(err)
	}
	err = myFirstAccount.Withdraw(30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myFirstAccount.Balance)
}
