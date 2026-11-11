# The Errors Package

The Go standard library provides an "errors" package that makes it easy to deal
with errors.

Tread the godoc for the error.New() function, but there's a simple example:

var err error = errors.New("something went wrong")

# Assignment

Textio software artchitects may have overcomplicated the requirements from the last coding 
assignment...oops. All we needed was a new generic error message that returns the string not dividing by 0
when a user attempts to get us to perform the taboo

Complete the divide function. If y == 0, return 0 and error created with errors.New() that reads "no dividing by 0."