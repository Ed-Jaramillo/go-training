package main

import "fmt"

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}
	return "", false
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {
	u := newUser("Alice", "premium")
	msg, ok := u.SendMessage("Hello, this is a test message!", 31)
	fmt.Println(msg, ok)
}