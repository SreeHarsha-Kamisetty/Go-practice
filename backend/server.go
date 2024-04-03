package main

import (
	"fmt"      // format
	"net/http" //http module
)

func homeRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home route")
}

func main(){
	http.HandleFunc("/",homeRoute)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}