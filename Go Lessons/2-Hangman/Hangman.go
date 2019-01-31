package hangman

import "fmt"
import "math/rand"
import "time"

func containsAny(s string, chars []string) bool {
	for _, ch := range chars {
		if string(ch) == ch2 {
			return true
		}
	}
	return false
}

func join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := ""
	lastIdx := len(strings) - 1
	for _, v := range strings[0:lastIdx] {
		s += v + separator
	}
	return s + strings[lastIdx]
}

func getLetter(found []string) string{
	alpahbet:="abcdefghijklmnopqrstuvwxyz"
	for true{
		letter:=prompt("Pick a letter",join(found," "))
	   if len(letter)=1 && containsAny (alphabet, []string(letters)) {
		   return letter
	   }
	   fmt.Println("Invalid input must enter a single  lowercase letter.")
	}
	return ""
}

func updateFound(found []string, word string,letter string) bool{
	 complete:=true 
	for i,r:=range word{
		if letter == string(r) {
			found[i]=letter
		}
		if found[i]=="_"{
			complete=false
		}
	}
	return complete
}

func main(){
words:=[]string{"zebra","moosse","aligator","elephant","ibex","jerboa","cat","hipopotamus","pterodaktyl"}
t:=time.Now()
rand.Seed(t.UnixNano())
idx:=rand.Intn(len(words))
word:=words[idx]
nGueses:=len(word)
found:=[] string{}
for i:=0;i<len(word);i++{
	found=append(found,"_")
}
for nGuesses>0{
	fmt.Println("You have",nGueses,"remaining guesses.")
	letter:=getLetter(found)
	if !containsAny(word,[]string(letter)){
		nGueses--
	}
	if updateFound(found,word,letter){
		fmt.Println("You win!The word was:",word)
		return
	}
}
fmt.Println("You lose!The word was:",word)
}


func prompt (vals ...interface{}) (string,error){
	if len(vals) !=0{
		fmt.Println(vals...)
	}
scanner:=buffio.NewScanner(os.Stdin)
scanner.Scan()
err:=scanner.Err()
if err !=nill{
	return "",err
}
return scanner.Text(), nil
}