package main

import (
	"fmt"
	myPackage "myawesomemodule/test006____interfaces"
)

func main() {
	myAccounts := []myPackage.IBankAccount{
		myPackage.NewComdirectBankAccount(),
		myPackage.NewSparkasseBankAccount(),
	}

	for _, acc := range myAccounts {
		acc.Deposit(1000)
		_ = acc.Withdraw(100)
		fmt.Printf("balance on %T: %.2f EUR\n", acc, float64(acc.GetBalance())/100.0)
	}

}
