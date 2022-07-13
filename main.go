package main

import (
	"fmt"
	"sync"
)

func main() {
	input := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			input <- num
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(input)
	}()
	out := make(chan int, 10)
	go func() {
		for i := range input {
			out <- i
		}
		close(out)
	}()
	for o := range out {
		fmt.Println(o)
	}
}
