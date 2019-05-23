package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting2(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting2(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb engtlishBot) getGreeting() string {
	return "Hi"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
