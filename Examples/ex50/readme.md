# The Error Interface

Go programs express errors with error values. An error is any that that implements
the simple built-in error interface:

type error interface {
	Error() string
}

When something can go wrong in a function, that function should return an Erroras its last return value. Any code that calls a function
that can return an error should handle errors by testing whether the error is nil.

#Atoi

Let's look at how the strconv.Atoi functino uses this pattern. The function signature of Atoi is:

func Atoi(s string) (int, error)

This means Atoi takes a string argument and returns two values: an integer and an error.
If the string can be successfully converted to an integer, Atoi returns the integer, Atoi returns the integer and a nil error.
If the conversion fails, it returns zero and a non-nil error.

Here is how would safely use Atoi:

