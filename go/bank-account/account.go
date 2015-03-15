package account

import "sync"

type Account struct {
	sync.Mutex
	open    bool
	balance int64
}

// Open(initalDeposit int64) *Account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{open: true, balance: initialDeposit}
}

// (Account) Close() (payout int64, ok bool)
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}

	a.open = false
	payout = a.balance
	a.balance = 0
	return payout, true
}

// (Account) Balance() (balance int64, ok bool)
func (a *Account) Balance() (balance int64, ok bool) {
	if a.open == false {
		return 0, false
	}
	return a.balance, true
}

// (Account) Deposit(amount uint64) (newBalance int64, ok bool)
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open || amount+a.balance < 0 {
		return a.balance, false

	}
	a.balance += amount
	return a.balance, true
}
