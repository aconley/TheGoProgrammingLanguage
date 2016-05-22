// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Modified for exercise 9.1: Add a withdraw function

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawAction struct {
	amount  int
	success chan<- bool
}

var withdrawls = make(chan withdrawAction) // send amount to withdraw
var deposits = make(chan int)              // send amount to deposit
var balances = make(chan int)              // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if amount < 0 {
		return false
	}
	if amount == 0 {
		return true
	}

	var suc = make(chan bool)
	defer close(suc)
	withdrawls <- withdrawAction{amount, suc}
	return <-suc
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wa := <-withdrawls:
			if wa.amount > balance {
				wa.success <- false
			} else {
				balance -= wa.amount
				wa.success <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
