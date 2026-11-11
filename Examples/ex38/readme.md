# Embedded Structs

Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times.
Keep in mind, Go doens't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields
between struct definitions.


type car {
	brand string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contais all
	// of the fields of the car struct
	car
	bedSize int
}

# Embedded vs. Nested

* Unlike nested structs, an embedded struct's fields are accessed
at the top level like normal field.

* Like nested structs, you assign the promoted fields with the embedded struct in a 
composite literal


lanesTruck := truck{
	bedSize: 10,
	car: car{
		brand: "Toyota",
		model: "Tundra",	
	},
}

fmt.Println(lanesTruck.brand) //Toyota
fmt.Println(lanesTruck.model) //Tundra