# Type Assertions in Go

When working with interfaces in Go, every once in-awhile you'll need access to the underlying type of an interface value.
You can cast an interface value. You can cast an interface to its underlying type using a type assertion.


The example below shows to safely access the radius field of s when it is not know that s is a circle:


type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	//....
}

func printShapeInfo(s shape) {
	c, ok := s.(circle)
	if ok {
		radius := c.radius
		fmt.Printf("s is a circle, radius: %v\n", radius)
	
	}
}

In printShapeInfo:

We want to check if s is a circle in order to cast it into its underlying 
concrete type.

* We know (from the function signature) that s is an instance of the shape interface,
but we do not know if it's also a circle

* c is a new circle struct cast from safely
* ok is true if is indeed a circle, or false if s is NOT a circle

c and ok aren't duplicates: the assertion returns the concrete circle value and a separate boolean that says whether the assertion succeeded.


