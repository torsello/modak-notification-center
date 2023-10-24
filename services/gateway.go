package services

import "log"

type Gateway struct{}

func (gateway Gateway) Send(user, message string) {
	log.Printf("sending message to user %s\n", user)
}