package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAmount = errors.New("amount must be greater than zero")

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
	fmt.Println(myFirstAccount.Balance)
}
