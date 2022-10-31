package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type myStruct struct {
	in  float64
	out float64
}

func doWork(n int) float64 {
	time.Sleep(time.Millisecond * 100)
	return math.Sqrt(float64(n))
}

func main() {

	ch := make(chan myStruct)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(1000)
		for i := 0; i < 1000; i++ {
			//			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				ch <- myStruct{
					in:  float64(i),
					out: doWork(i),
				}
			}(i)
		}
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Printf("the square root of %v is %v\n", x.in, x.out)
	}

}
