package main

import "fmt"

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
	}
}

func (a *Account) WithDraw(amount float64) bool {
	if amount > 0 && a.Balance >= amount {
		a.Balance -= amount
	}

	return true
}

func (a Account) GetBalance() float64 {
	return a.Balance
}

func main() {
	acc := Account{Balance: 1000}
	fmt.Println("Current Balance:", acc.GetBalance())
	acc.Deposit(100)
	fmt.Println("Balance after deopsit", acc.GetBalance())
	acc.WithDraw(900)
	fmt.Println("Balance after withdraw", acc.GetBalance())
}
