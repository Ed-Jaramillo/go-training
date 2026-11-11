# Message Formatter

As Textio evolves, the team has deciddd to introduce a new feature for custom
message formats. Depending on the user's preferences, messages can be sent in different 
formats, such as plain text, markdown, code, or even encrypted text. To efficiently manage this, you'll implement a system using interfaces.


# Assignment

1. Implement formatter interface with a method format that returns a formatted string.
2. Define structs that satisfy the formatter interface: plainText, bold, code.
 * The structs must all have a message field of type string
3. plainText should return the message as is.
4. bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **mesage**).
5. code should wrap the message in a single backtick () to simulate inline code (e.g., `message`)


