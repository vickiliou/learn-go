package pointers

import "fmt"

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

// String defines how Bitcoin type is printed
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds Bitcoin to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the total of Bitcoin that a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
