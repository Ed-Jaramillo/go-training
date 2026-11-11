# Struct Methods in Go

While go is not objec-oriented, it does support methods that can be defined on structs. Methdods are just functinos that have a receiver. 
A receiver is a special parameter that syntactically goes before the name of the function.


type rect struct {
	width int
	height int

}

// area has a receiver of (r rect)
// rect is the Struct
// r is the placeholder


func (r rect) area() int {
	return r.width * r.height
}

var r = rect{
	width: 5,
	height: 10,
}

fmt.Println(r.area())