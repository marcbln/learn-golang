package test006____interfaces

type IBankAccount interface {
	GetBalance() int // 100 is 1 EUR
	Deposit(amount int)
	Withdraw(amount int) error // if not enough money on bank account it returns error
}
