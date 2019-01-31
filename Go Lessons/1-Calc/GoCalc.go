package main

import "fmt"

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

func main() {
	numbers := []int{6, 2, 1, 8, -56}
	result := sum(numbers)
	fmt.Println(result)
}
