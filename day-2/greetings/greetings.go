package greetings

import (
	"errors"
	"fmt"
)


func Hello(name string) (string,error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello %v. How are you?",name)

	// fmt.Println(message)
	return message,nil
}