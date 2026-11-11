# Functions

Functions in Go can take zero or more arguments.

The types of those arguments (or parameters) need to be declared explicitly. You put the type 
after the parameter name.

For example, the following function:

func sub(x int, y int) int {
	return x-y
}

Accepts two intefer parameters, x and y, and returns another integer.

The line func sub(x int, y int) int is known as the function signature.

Go doesn't care where in a file function is defined: you can call a function that's 
declared later, even below main.

# Assignment

We'll often need to manipulate strings in our messaging app - for example, 
adding some personalization by using a customer's name within a template. The concat
function should take two strings and smash them together.

"hello" + "world" = "helloworld"