package main

import "fmt"

type messageToSend struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	
	if mToSend.recipient.name == "" {
		return false
	}
	
	if mToSend.recipient.number == 0 {
		return false		
	}
	return true
}

func main() {
	m := messageToSend{
		message:   "Hello!",
		sender:    user{name: "Alice", number: 1},
		recipient: user{name: "Bob", number: 2},
	}
	fmt.Println("Can send message:", canSendMessage(m))
}