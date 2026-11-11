# Ignoring Return Values

A function can return a value that the caller doesn't care about. We
can explicitely ignore variables by using an underscore, or more precisely, the blank_identifyer _

for example:

func getPoint() (x int, y int) {
	return 3, 4
}

// ignore y value

x, _ := getPoint()

# Assignment

1. Run the code as is. You should get a compiler error.
2. Fix getProductMessage to ignore the unused return value

