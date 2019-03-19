package main
import (
"fmt"
)

func sortedKeys(m map[Key] Value)([]Key){
keys:=make([]Key, len(m))
i:=0
for k:=range m{
	keys[i]=k
	i++
}
sort.Keys(keys)
return keys


}
 func main() {
var romanNumeralDict map[int]string = map[int]string{
  1000: "M",
  900 : "CM",
  500 : "D",
  400 : "CD",
  100 : "C",
  90  : "XC",
  50  : "L",
  40  : "XL",
  10  : "X",
  9   : "IX",
  5   : "V",
  4   : "IV",
  1   : "I",
}

sortedKeys(romanNumeralDict)
for _,k:=range keys{
	fmt.Println(k,romanNumeralDict[k])
}
	
}