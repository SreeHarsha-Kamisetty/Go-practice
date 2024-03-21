package main

import (
	"fmt"
	"greetings"
)

func main(){
	message := greetings.Hello("Harsha")

	fmt.Println(message)
}