#Anonymous Functions

Anonymous functions are true to form in that they have no name.
They're useful when defining a function that will only be used
once or to create a quick closure.

Let's say we have a function conversions that accepts another function, converter as input:


func conversions(converter func(int) int, x, y, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ	
}

We could define a function normally and then pass it in by name....
but it's usually easier to just define it anonymously:


func double(a int) int {
	return a + a
}

func main() {
	// using a named function
	newX, newY, newZ := conversions(double, 1,2,3)
	// newX is 2, newY is 4, newZ is 6
	
	// using an anonymous function
	newX, newY, newZ = conversions(func(a int) int {
	return a + a
	
	}, 1,2,3)
	
}

#Assignment 

Complete the printReports function. It takes as input a sequene of messages,
intro, body, outro. It should call printCostReport once for each message
by making three separate calls (you don't need to use loops or arrays for this).

For each call of printCostReport, give it an anonymous function that returns the cost
of a message as an integer. Here are the costs.

* Intro: 2. the message length
* Body: 3x the message length
* Outro 4x the message length