package main

import (
	"fmt"
	"time"
)

func main() {
	//process("order")
	go process("order")
	process("refund")
}

func process(item string) {
	for i := 1; true; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
	}
}
