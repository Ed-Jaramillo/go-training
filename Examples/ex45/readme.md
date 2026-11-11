# Multiple Interfaces

A type can implement any number of interfaces in Go. For
example, the empty interface, interface{}, is always implemented by
every type because it has no requirements.


# Assignment 

Complete the required methods so that the email type implements both the expense and formatter interfaces.

Complete the cost() method:

* If the email is not "subscribed", then the cost is 5 cents
* If it is, then the cost is 2 cents per character
* Return the total cost of the entire email in cents
Complete the format() method.

It should return a string in this format:

'CONTENT' | Subscribed

If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed