package main

import "fmt"

type MyStruct struct {
	a int
	b int
}

func (ms MyStruct) String() string {
	return fmt.Sprintf("(%d,%d)", ms.a, ms.b)
}

func (ms MyStruct) myFuncWithValueReceiver(a, b int) {
	ms.a = a
	ms.b = b
}

func (ms *MyStruct) myFuncWithPointerReceiver(a, b int) {
	ms.a = a
	ms.b = b
}

func main() {
	ms := MyStruct{
		a: 0,
		b: 0,
	}
	fmt.Println(ms) // (0,0)

	ms.myFuncWithValueReceiver(33, 44)
	fmt.Println(ms) // (0,0)

	ms.myFuncWithPointerReceiver(55, 66)
	fmt.Println(ms) // (55,66)

}
