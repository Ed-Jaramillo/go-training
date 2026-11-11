package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.54
	
	// don't edit above this line
	
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent \n", name, openRate)
	
	// don't edit below this line
	fmt.Print(msg)
}