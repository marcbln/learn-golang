package test006____interfaces

import "errors"

type SparkasseBankAccount struct {
	balance int
	fee     int
}

// constructor
func NewSparkasseBankAccount() *SparkasseBankAccount {
	return &SparkasseBankAccount{
		balance: 0,
		fee:     250,
	}
}

func (s *SparkasseBankAccount) GetBalance() int {
	return s.balance
}

func (s *SparkasseBankAccount) Deposit(amount int) {
	s.balance += amount
}

func (s *SparkasseBankAccount) Withdraw(amount int) error {
	subtrahend := amount + s.fee
	if amount < subtrahend {
		s.balance -= subtrahend
		return nil
	} else {
		return errors.New("insufficient funds")
	}
}
