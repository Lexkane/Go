package main


import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w,"My Awesome Go App")	
}
func setupRoutes(){
	http.HandleFunc("/",homePage)
}

func main(){
	fmt.Println("GO web abb started on port 3000")
	setupRoutes()
	http.ListenAndServe(":3000",nil)
}