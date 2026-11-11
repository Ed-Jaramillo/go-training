# Why Go?

* It's fast. Go compiles to a single native binary, and it compiles fast.

* It's simple. It's a small language without a ton of feature bloat.

* It's statistically typed. The compiler catches type bugs
before your code even runs

* Concurrency is built in. Goroutines are very powerful, which is a big reason tools like 
Docker and Kubernetes are written in Go

# Assignment 

## Critical bug!

Textio users reported that we're billing them for wildly inaccurate amounts. They're
supposed to be billed .02 dollars (2 cents) for each text message sent.


Without changing any of the other lines, *fix the math bug in the totalCost calculation*