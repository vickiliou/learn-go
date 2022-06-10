package pointers

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

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
