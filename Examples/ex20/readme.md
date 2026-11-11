#Conditionals

if statemets in Go do not use parentheses around the condition:

if height > 4 {
	fmt.Println("You are tall enough!")

}

else if and else are supported as you might expect:

if height > 6 {
	fmt.Println("You are super tall!")
} else if height > 4 {
	fmt.Println("You are tall enough!")
} else {
	fmt.Println("You are not tall enough!")
}

#Assignment 
Fix the bug on line 12. If messageLen is less than or equal to the 
maxMessageLen the program should print "Message sent", else it 
should print "Message not sent"