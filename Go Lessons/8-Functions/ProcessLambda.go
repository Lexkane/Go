package main

import "fmt"

func process(nums []int ,f func (int) int ){
	for i,v:=range nums{
		nums[i]=f(v)
	}
}

func main(){
	val:=promptInt()
	x:=[]int {3,2,4,-10,8}
	process{x,func(a int ) int{
		retunr a+val
	})
	fmt.Println(x)
}