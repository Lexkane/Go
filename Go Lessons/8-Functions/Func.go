package main


type cat struct{
name string
age int
weight float32
}

func (c cat) eat (food float32) float32{
	return c.weight+food
}

func (c cat) eats (food float32) int{
	//
}

func main(){
c:=cat{"Oscar",14,15.2}
c2:=cat{"Mittens",10,8.4}
var f func (float32) float32
f=c.eat
f(0.4)
f=c2.eat
var foo func(cat,float32) int
var bar func (*cat, float32) int
foo=cat.eats
bar=(*cat).eats
var p cat
foo (p,1.7)
bar(&p,1.7)

d:=func(a int) int {
	return a+3
}




}