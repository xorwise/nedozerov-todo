package services

import (
	"github.com/xorwise/nedozerov-todo/internal/domain"
	"sync"
)

type account struct {
	id      int
	balance float64
	mu      *sync.Mutex
}

func NewAccount(id int) domain.BankAccount {
	return &account{
		id:      id,
		balance: 0,
		mu:      &sync.Mutex{},
	}
}

func (a *account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return domain.ErrInvalidAmount
	}
	a.balance += amount
	return nil
}

func (a *account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return domain.ErrInvalidAmount
	}
	if a.balance < amount {
		return domain.ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

func (a *account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
