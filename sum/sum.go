package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	c := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			sum++
			c <- sum
			wg.Done()
		}()
		sum = <-c

	}
	wg.Wait()
	fmt.Println(sum)
}
