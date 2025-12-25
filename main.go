package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("Deposited: %v, New Balance: %v \n ", amount, b.Balance)
}
func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance >= amount {
		b.Balance -= amount
		fmt.Printf("Withdrawn: %v, New Balance: %v \n", amount, b.Balance)
	} else {
		fmt.Println("❌ Insufficient Funds!")
	}
}
func (b BankAccount) PrintStatement() {
	fmt.Printf("Name: %v, Balance: %v \n ", b.Owner, b.Balance)
}

func main() {
	bankAccount := BankAccount{
		Owner:   "Yasser",
		Balance: 0.0,
	}
	bankAccount.Deposit(1500)
	bankAccount.Withdraw(500)
	bankAccount.PrintStatement()
	bankAccount.Deposit(600)

}
