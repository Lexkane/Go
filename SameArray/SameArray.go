package main

import "fmt"

func main() {

	arr1 := []string{"aa", "bb", "cc"}
	arr2 := []string{"aa", "bb", "cc"}
	fmt.Println(checker(arr1, arr2))
}

func checker(ary1 []string, ary2 []string) bool {
	same := true
	for i, elm := range ary1 {
		if ary2[i] != elm {
			same = false
		}
	}

	return same
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
