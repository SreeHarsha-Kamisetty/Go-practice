package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string)(string,error){

	if name == "" {
		return "", errors.New("please provide a name")
	}

	// message := fmt.Sprintf("Hi %s. How was your day ? :)",name)

	message := fmt.Sprintf(randomFormat(),name)

	return message,nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %s. Welcome !",
		"Great to see you, %s!",
		"Hey!!!!!!!, %s",
	}

	return formats[rand.Intn(len(formats))]
}