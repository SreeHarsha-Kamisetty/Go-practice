package main

import (
	"fmt"      // format
	"net/http" //http module
)

func homeRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home route")
}

func aboutRoute(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"About route")
}

func contactRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Contact route")
}
func main(){
	http.HandleFunc("/",homeRoute)
	http.HandleFunc("/about",aboutRoute)
	http.HandleFunc("/contact",contactRoute)
	
	
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}