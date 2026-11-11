#Runes and String Encoding

In many programming languages (C), a "character" is a single byte.
Using ASCII enconding, the standard for the C programming language, we can 
represent 128 characters.

In Go, strings are just sequences of bytes: they can hold arbitrary data. However,
Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough
to hold any Unicode code point.

When you're working with strings, you need to be aware of the encoding (bytes -> representation). Gouses UTF-8 encoding,
which is a variable length encoding.

UTF-8 Text

Type text to see its code points and bytes.


# Assignment

Boots is a bear.

1. Run the code as-is. Notice that the simple string "boots" has 5
bytes, and 5 runes (characters).

2. Update the 'name' constant to be the bear emoji instead of the
word "boots".

