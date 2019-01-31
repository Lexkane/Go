package main

import (
	"bufio"
	"fmt"
	"os"
)

func FibonacciGenerator(first uint64, second uint64) chan uint64 {
	c := make(chan uint64)

	go func() {
		for i, j := first, second; ; j, i = i+j, j {
			c <- i
		}
	}()

	return c
}

func main() {
	generator := FibonacciGenerator(1, 1)
	for i := 0; i < 150; i++ {
		if _<-generat
		fmt.Println(<-generator)
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
