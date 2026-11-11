# Short Variable Declaration

Sad variable declaration:

var mySkillIssues int
mySkillIssues = 42

GOATed variable declaration:
mySkillIssues := 42


The walrus operation, :=, declares a new variable and
assigns a value to it in one line. Go can infer that
mySkillIssues is an int because of the 42 value. Yay type inference

# When to Use the Walrus Operator
The := (walrus operator) should be used instead of var
style declarations basically anywhere possible. The limitation is that 
:= can't be used outside a function 

Type inference is based on the value being assigned.

an int:

mySkillIssues := 42

A float64:

pi := 3.14159

A string:

message := "Hello, world!"

A bool:

isGoat := true

# Assignment 

A common use case for Textio is to send birthday messages.

1. Complete the main function. It should print "Happy birthday! You are now 21 years old!".
1.1 Create a string variable messageStart with the text "Happy birthday! You are now"
1.2 Create an integer variable age set to 21
1.3 Create another string variable messageEnd with the text "years old!"


2. The provided fmt.Println statemente will print the full message on a single line separated by spaces.

