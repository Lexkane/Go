package main

import "fmt"
func factorial (n int) int {
	if n<0 {
		panic("cannot compute factorial of negative")
	}
	if n==0{
		return 1
	}
	return n*factorial(n-1)
}

func cat(){
	fmt.Println("meow")
}

func dog(a int){
	fmt.Printtln("dog",a)
}

func bar(){
	defer dog()
	panic("bad dog")
}

func foo(){
	defer cat()
	bar()
}
func main(){
	defer cat()
	i:=10
	efer dog(i)
	i=6
	defer dog (i)

	foo()
	fmt.Println("end main")
}