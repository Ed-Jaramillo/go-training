package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}

type bold struct {
	message string
}

type code struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

func (b bold) format () string {
	return "**" + b.message + "**"
}

func (c code) format () string {
	return "`" + c.message + "`"
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}

func main() {
	messages := []formatter{
		plainText{message: "Hello there!"},
		bold{message: "Hello there!"},
		code{message: "Hello there!"},
	}

	for _, m := range messages {
		fmt.Println(sendMessage(m))
	}
}

