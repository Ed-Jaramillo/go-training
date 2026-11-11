# Constants

Constants are declared with the const keyword. They can't use the := short 
declaration syntax.

const pi = 3.14159

Constants can be primitive types like strings, integers, booleans and floats.
They cannot be more complex types like slices, maps and structs, which are types we will explain later.


As the name implies, the value of a constant can't be changes after it has been declared



# Use Two Separate Constants

Something weird is happening in this code.

What should be happening is that we create 2 separate constants:
premiumPlanName and basicPlanName. Right now it looks like we're trying to overwrite one of them.

# Assignment 

Complete the code to remove the bug and create the constant basicPlanName