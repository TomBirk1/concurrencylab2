package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	c := make(chan int, 1)
	total := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			sum++
			c <- sum
			wg.Done()
		}()
		total = <-c

	}
	wg.Wait()
	fmt.Println(total)
}
