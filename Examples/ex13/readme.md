# Computed Constants

Constants must be known at compile time. They are usually declared with a static value.

const myInt = 15

However, constants can be computed as long as the computation can happen at compile time. 

For example, this is valid:

const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time
like you can in JavaScript. This breaks:

// the current time can only be known when the program is running

const currentTime = time.Now()