# User Input

In Textio, users can set their profile status to communicate their current activity to
those that choose to read it.... However, there are some restrictions on what
these statuses can contain. Your task is to implement a function that validates a user's status update. the
status update cannot be empty and must not exceed 140 charactesr.

# Assignment

Complete the validateStatus function. It should return an error when the status update violates any of the rules

* If the status is empty, retunrn an error that reads status cannnot be empty.
* If the status exceeds 140 characters, retrn an error that says status exceeds 140 characers
* Otherwise, return nil.