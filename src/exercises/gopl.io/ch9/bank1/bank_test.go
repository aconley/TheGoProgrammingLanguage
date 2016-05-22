package bank_test

import (
	"testing"

	"exercises/gopl.io/ch9/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Now withdraw

	// Try to withdraw too much
	if bank.Withdraw(310) {
		t.Errorf("Should not have been able to withdraw 310")
	}

	if !bank.Withdraw(10) {
		t.Errorf("Should have been able to withdraw 10")
	}
	if got, want := bank.Balance(), 290; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
