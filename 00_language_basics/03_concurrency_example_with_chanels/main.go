package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			out1 <- "order processed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "refund processed"
		}
	}()

	for {
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
		fmt.Println(<-out1)
		fmt.Println(<-out2)
	}
}
