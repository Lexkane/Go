package main

import "fmt"

func cat(){
	fmt.Println("meow")
	v:=recover()
	fmt.Println(v)
}

func dog(a int){
	fmt.Println("bark")
}



func foo() int {
	defer cat()
	defer dog()
	panic("abort")
	fmt.Println("end foo")
	return 5
}

func main(){
	z:=foo()
	fmt.Println("after",z)
	fmt.Println("end main")
}