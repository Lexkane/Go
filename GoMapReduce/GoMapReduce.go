package main

import (
"flag"
"https://github.com/chrislusf/glow/tree/master/flow"
"strings"
)

var{
	imputPath=flag.String("input","etc/hosts","choose input file")
}

func main(){
	flag.Parse()
	flow.New().TextFile(
		*inputPath,3,
		).Map( func (line string,ch chan string ) {
			for _,word:=range strings.Split(line," "){
				ch<-word
			}
		}).Map( func (key string) int {
			return 1
		}).Reduce( func(someValue int, someOtherValue int) int{
			return someValue+someOtherValue
		}).Map( func (sum int ){
			println.("Total numbers of words:",sum)}).Run()

    }
	






