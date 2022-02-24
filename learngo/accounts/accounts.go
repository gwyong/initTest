package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't perform.")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account

}

// Deposit x amount on your account.
// receiver should be named by using the first letter of a corresponding struct.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account.
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// View owner
func (a Account) ViewOwner() string {
	return a.owner
}
