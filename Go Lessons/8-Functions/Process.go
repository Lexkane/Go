package  main

func process(num [] int, f func(int) int){
	for i,v:=range nums{
		nums[i]=f(v)
	}
}

func adder (a int ) int {
	return a+5
}

func main(){
	x:=[] int {3,2,4,-10,8}
	process(x, func (a int ) int {
		return a+5
	})
	fmt.Println(x)
}