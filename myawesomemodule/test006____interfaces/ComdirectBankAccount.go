package test006____interfaces

import "errors"

type ComdirectBankAccount struct {
	balance int
}

// constructor
func NewComdirectBankAccount() *ComdirectBankAccount {
	return &ComdirectBankAccount{
		balance: 0,
	}
}

func (s *ComdirectBankAccount) GetBalance() int {
	return s.balance
}

func (s *ComdirectBankAccount) Deposit(amount int) {
	s.balance += amount
}

func (s *ComdirectBankAccount) Withdraw(amount int) error {
	if amount < s.balance {
		s.balance -= amount
		return nil
	} else {
		return errors.New("insufficient funds")
	}
}
