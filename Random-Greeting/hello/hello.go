package main

import (
	"fmt"
	"log"

	"greetings"
)

func main(){

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message,err := greetings.Hello("Sree Harsha")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}