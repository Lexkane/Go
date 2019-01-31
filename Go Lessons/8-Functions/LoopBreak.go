package  main

func main(){
  loop:
for i:=0;i<10;i++{
	for j:=0;j<5;j++{
		if someCondition{
			break loop
		}
		doSomething()
	}
}

}
