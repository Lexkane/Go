package main

 import "fmt"


 func ack(){
	 recover()
 }

 func bar(){
	 ack()
	 recover()
}

func foo(){
	defer bar()
	panic("abort")
}

func foo2(){
	defer func(){
		recover
	}()
	panic("abort")
}

func foo3 (a int )( x int ){
	x=5+a
	if x>10{
		return 
	}else{
		return 10
	}
}

func bar2( a int) ( x int, y int){
	x=5+a
	y=5*a
	if x>10{
		return
	} else{
		return 10,y
	}
}


func main(){
	foo()
	foo2()
	fmt.Println("endmain")
}
