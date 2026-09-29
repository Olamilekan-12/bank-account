package main

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name        string
		amount      int
		wantErr     error
		wantBalance int
	}{
		{"valid deposit", 100, nil, 100},
		{"zero amount", 0, ErrInvalidAmount, 0},
		{"negative amount", -50, ErrInvalidAmount, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myAccount := &Account{}
			err := myAccount.Deposit(tt.amount)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Deposit(%d) error = %v, want %v\n", tt.amount, err, tt.wantErr)
			}
			if myAccount.Balance != tt.wantBalance {
				t.Errorf("unexpected account balance return. Expected : %v, got: %v", tt.wantBalance, myAccount.Balance)
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name             string
		startBalance     int
		amount           int
		wantErr          error // sentinel errors, checked with errors.Is
		wantInsufficient bool  // custom type, checked with errors.As
		wantBalance      int
	}{
		{"valid withdrawal", 100, 30, nil, false, 70},
		{"zero amount", 100, 0, ErrInvalidAmount, false, 100},
		{"exceeds balance", 100, 500, nil, true, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myAccount := &Account{
				Balance: tt.startBalance,
			}
			var insufficient *InsufficientFundsError
			err := myAccount.Withdraw(tt.amount)
			if !tt.wantInsufficient && !errors.Is(err, tt.wantErr) {
				t.Errorf("[Invalid amount] Withdraw(%d) error(%v), want %v\n", tt.amount, err, tt.wantErr)
			}
			if errors.As(err, &insufficient) != tt.wantInsufficient {
				t.Errorf("[Invalid insufficient fund error] Withdraw(%d) error(%v)\n", tt.amount, err)
			}
			if myAccount.Balance != tt.wantBalance {
				t.Errorf("[Invalid balance] Start balance:(%d), Withdraw:(%d), Balance: (%d),  Wantbalance:(%d)", tt.startBalance, tt.amount, myAccount.Balance, tt.wantBalance)

			}
		})
	}
}
