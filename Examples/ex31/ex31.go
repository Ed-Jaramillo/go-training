package main

import "fmt"

func splitEmail(email string) (string, string) {
	username, domain := "" ,""
	
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func main() {
	email := "john.doe@example.com"
	username, domain := splitEmail(email)
	fmt.Println("Username:", username)
	fmt.Println("Domain:", domain)
}