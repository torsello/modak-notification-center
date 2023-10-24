package services

import "fmt"

type Gateway struct{}

func (gateway Gateway) Send(user, message string) {
	fmt.Printf("sending message to user %s\n", user)
}