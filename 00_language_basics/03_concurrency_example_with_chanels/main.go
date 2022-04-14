package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	go process("order", out1)
	for {
		msg, open := <-out1
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func process(item string, out chan string) {
	defer close(out)
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second / 2)
		//fmt.Println("Processed", i, item)
		out <- item
	}
	//close(out)
}
