package main

import (
	"fmt"
)

type cat struct{
name string
age int
weight float32
}

func (c cat) sleep (){
	fmt.Println("zzzz")
}

func (c *cat)eat (food float32) float32{
	c.weight+=food
	return c.weight+food
}

func main() {
	c:=cat ("Oscar",14,15.2)
	c.sleep(0)
	c.eat(0.3)
	p:=&c
	p.sleep()
	p.eat(0.3)
}