package main

import "fmt"

type bot interface {
	printGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting2(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb engtlishBot) getGreeting() string {
	return "Hi"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
