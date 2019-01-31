package main2

import (
	"fmt"
)

type cat struct{
name string
age int
weight float32
}

type anumal interface {
	eat()
	sleep()
	drink()
}

func (c *cat)  eat() {}
func (c cat) sleep(){}
func (c cat) drink() {}

func main(){
	var a cat
	a=cat{}
	b:=&cat{}
    fmt.Println(a)
}