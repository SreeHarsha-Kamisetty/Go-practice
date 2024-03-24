package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handlerFunc)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)

}

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello , World!")
}