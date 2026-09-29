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

type InsufficientFundsError struct {
	Requested int
	Balance   int
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds: Requested %d, Balance %d", e.Requested, e.Balance)
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
		return &InsufficientFundsError{
			Requested: amount,
			Balance:   a.Balance,
		}
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
	var insufficient *InsufficientFundsError
	if errors.As(err, &insufficient) {
		fmt.Printf("Your account is short %d\n", insufficient.Requested-insufficient.Balance)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myFirstAccount.Balance)
}
