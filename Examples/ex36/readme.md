# Structs in Go

We use structs in Go to represent structured data. I'ts often convenient to group
different types of variables together. For example, if we want to represent a car we could do the following:

type car struct {
	brand string
	model string
	doors int
	mileage int
}

This creates a new struct type called car. All cars have a brand, model, doors an mileage.

To create a car, use struct literal:

myCar := car{
	brand: "Toyota",
	model: "Camry",
	doors: 4
	mileage: 5000
}

Structs in Go are often used to represent data that you might use a dictionary
or object for in other languages.

# Assignment
Complete  the definition of the messageToSend struct. It needs two fields:


* phoneNumber - an integer
* message - a string.