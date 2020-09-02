package wallet

import "fmt"

// Wallet - contains a balance
type Wallet struct {
	balance int
}

// Deposit - puts money into a wallet
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance - retrieves the balance of a wallet
func (w *Wallet) Balance() int {
	return w.balance
}
