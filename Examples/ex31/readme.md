# Block Scope

Unlike Python, Go is not function-scoped, it's block-scoped. Variables declared inside
a block are only accessible within that block (and its nested blocks). There's also the package scope.
We'll talk about packages later, but for now, you can think of it as the outermost, nearly global scope.


package main

// socped to the entire "main" package (basically global)

var age = 19

func sendEmail() {
	// scoped to the "sendEmail" function
	name := "John Snow"
	
	for i := 0; i < 5; i++ {
	// scoped to the "for" body
	
	email := "snow@winterfell.net"
	
	}

}


Blocks are defined by curly braces {}. New blocks are created for:

* Functions
* Loops
* If statements
* Switch statements
* Select statements
* Explicit blocks

