# Memory Layout

In Go, structs sit in memory in a contiguous block, with fields
placed one after another as defined in the struct. For example
this struct.

type stats struct {
	Reach uint16
	NumPosts uint8
	NumLikes uint8
}


# Field ordering....Matters?

the order of fields in a struct can have big impact on memory usage. THi is the same struct as above, but poorly designed.

the order of fields  in a struct can have big impact on memory usage. This is the same struct as above, but poorly designed:


type stats struct {
	NumPosts uint8
	Reach uint16
	NumLikes uint8
}

# Empty struct
Empty structs are used in Go as a unary value

// anonymous empty struct type

empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}