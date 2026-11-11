# Functions As Values

Go supports first-class and higher-order functions, which are just fancy
ways of saying "functions as values"

Functions are just another type -- like int s and string s and bool s
Let's assume we have two simple functions:

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

We can create a new aggregate function that accepts a function as it's 4th argument:

func aggregate (a, b, c int, arithmetic func(int, int) int ) int {
	firstResult := arithmetic(a,b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}

func main () {
	sum := aggregate(2,3,4,add)
	// sum is 9
	product := aggregate(2,3,4, mul)
	// product is 24
}


# Assignment 

Complete the reformat function. It takes a message string and a formatter function as input.

1. Apply the given formatter three times to the message
2. Add a prefix of TEXTIO: to the result
3. Return the final string

For example, if the message is "General Kenobi" and the given formatter adds a period to 
the end of the string, the final result should be

TEXTIO: General Kenobi......


