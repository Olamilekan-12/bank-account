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
