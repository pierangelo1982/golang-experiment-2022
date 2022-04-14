package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // number of process

	go func() {
		process("refund")
		wg.Done()
	}()

	go func() {
		process("order")
		wg.Done()
	}()

	wg.Wait()
}

func process(item string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
	}
}
