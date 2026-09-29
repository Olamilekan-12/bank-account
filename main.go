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
		return ErrInsufficientFunds
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
	err = myFirstAccount.Deposit(-50)
	if err != nil {
		fmt.Println(err)
	}
	err = myFirstAccount.Withdraw(500)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Printf("Not enough money. Current balance: %d\n", myFirstAccount.Balance)
	} else if err != nil {
		fmt.Println(err)
	}
	err = myFirstAccount.Withdraw(-10)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Printf("Not enough money. Current balance: %d\n", myFirstAccount.Balance)
	} else if err != nil {
		fmt.Println(err)
	}

	err = myFirstAccount.Withdraw(30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myFirstAccount.Balance)
}
