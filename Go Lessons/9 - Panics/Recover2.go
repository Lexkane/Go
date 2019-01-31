package main

import "fmt"


func  foo( x int){
	defer func(){
		if recover()==nil{
			x=5
		}
	}()
	return
}

func main(){
	z:=foo()
	fmt.Println("end main",z)
}