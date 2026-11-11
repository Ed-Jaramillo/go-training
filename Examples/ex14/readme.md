# Formatting Strings in Go

Go follows the printf tradition from the C language. In my opinion,
string formatting/interpolation in Go is less elegant than Python's 
f-strings, unfortunately.

fmt.Printf() - Prints a formatted string to standard out.
fmt.Sprintf() - Returns the formatted string

These following "formative verbs" work with the formatting functions above

s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old

If you want to print in a more specific way, you can use the following formatting verbs:

# string
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old.


# Integer
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

#Float 

s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old

If your're interested in all the formating options, you can look at the fmt package's doc

# Assignment

Create a new variable called msg on line 11 and use the approrpiate formatting function to return a string
that contains the following.

"Hi NAME, your open rate is OPENRATE percentNEWLINE"

* Replace NAME with the variable name,
* Replace OPENRATE with the variable openRate rounded to the nearest "tenths" place, e.g 10.523
should be rounded down to 10.5
* The word percent should apppear as part of the strging following the open rate value
* Replace NEWLINE with the \n escape sequence.

For example, with the inputs "Jimmy McGill" and 2.5, the expected output would be"

'Hi Jimmy McGill, your open rate is 2.5 percent'
