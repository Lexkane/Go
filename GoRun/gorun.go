package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	x := []int{2}

	for i := 3; i < 200000; i += 2 {
		simple := true

		for _, j := range x {
			if i%j == 0 {
				simple = false
				break
			}
		}

		if simple {
			x = append(x, i)
		}
	}

	fmt.Println(x)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
