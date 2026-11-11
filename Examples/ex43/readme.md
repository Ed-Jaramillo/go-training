# Interfaces in Go

Interfaces allow you to focus on what a type does rather than how it's built. They can help you
write more flexible and reusable code by defining behaviors (like methods) that different types can share
This makes it easy to  swap out or update parts of your code without changing everything else.


Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the 
interface if it has methods that amtch the interface method signatures.

In the following example, a "shape" must be able to return its area and perimeter. Both rect and circle fulfill the interface.

type shape interface {
	area () float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}


func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
}

Because we say the input is of type shape, we know that any argument must implement the .area() and .perimeter() methods.