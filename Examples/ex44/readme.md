# Interface Implementation

Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has
the proper methods defined, then the type automatically fulfills that interface.

A quick way of checking whether a struct implements an Interface
is to declare a function that takes an interface as an argument.
If the function can take the struct as an argument, then the struct implements the interface.

# Assignment 

At Textio we have full-time employees and contract employees. We have been tasked with making a more general employee
interface so that dealing with different employee types is simpler.



