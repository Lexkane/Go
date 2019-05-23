package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func RandomSlice(items []int, number int) ([] string) {
	list := make([]int, 0)
	newlist := make([]string, 0)
	list = rand.Perm(number)
	for i, _ := range list {
		list[i]++
	}
	for i, _ := range list {
		s, _ := json.Marshal(list[i])
		newlist = append(newlist, string(s))
	}
	return newlist
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	List := make([]int, 0)
	var StringList = RandomSlice(List, 50)
	fmt.Println("List before reverse")
	fmt.Println(StringList)
	ReversedList := make([]string, 0)
	/*
		scanner:=bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			s=scanner.Text()
		List=append(List, s)
		}
	*/
	for i, _ := range StringList {
		s := Reverse(StringList[i]);
		ReversedList = append(ReversedList, s)
	}
	fmt.Println("List after reverse")
	fmt.Println(ReversedList)

}
