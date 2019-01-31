package main
import {
"github.com/chrislusf/glow/flow"
"strings"
}

func main(){
count:=0
var lastword string=""

flow.New().TextFile("Gamelocalisation",5,).Partition(
2,
).Filter(func (line string) bool{
return !strings.HasPrefix(line,"***:)
}).Map(func(line string, ch chan string){
for _,word:=range strings.Split(line,""){
ch <-word
}
}).Sort(func (a string, b string ) bool{
return strings.Compare(a,b)<0
}).Reduce (func(a string, b string) string{
if lastWord==""{
lastWord=a
count=1
}
if b==lastWord{
count++
} else{
	println(lastWord,":",count)
	lastWord=b
	count=1
	}
	return "#NextWord"
}).Run()
}