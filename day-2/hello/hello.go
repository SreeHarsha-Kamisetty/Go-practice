package main

import (
	"fmt"
	"greetings"
	"log"
)

func main(){

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Harsha")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}