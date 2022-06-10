package pointers

// Wallet stores the number of Bitcoin.
type Wallet struct {
	balance int
}

// Deposit adds Bitcoin to the wallet.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns the total of Bitcoin that a wallet has.
func (w *Wallet) Balance() int {
	return w.balance
}
