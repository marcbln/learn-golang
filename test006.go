package main

import (
	"fmt"
	myPackage "myawesomemodule/test006____interfaces"
)

func main() {
	acc := myPackage.NewSparkasseBankAccount()
	acc.Deposit(100)
	err := acc.Withdraw(60)
	if err != nil {
		panic(err)
	}
	fmt.Printf("my account balance: %v\n", acc.GetBalance())

	err = acc.Withdraw(60)
	if err != nil {
		panic(err)
	}

}
