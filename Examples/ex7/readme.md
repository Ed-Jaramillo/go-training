# The Compilation Process

Computers need machine code, they don't understand English or even Go. We
need to convert our high-level (Go) code into machine language, which is really
just a set of instructions that some specific hardware can understand. 

The Go compiler's job is to take Go code and produce machine code, an
.exe file on Windows or a standard executable on Mac/Linux



# Two Kinds of Errors

Loosely speaking, there are two kinds of erros in programing:

1. Compilation errors happen when the code is compiled. They're the good kind: a program 
with a compiler error won't build, so it can't ship to production

2. Runtime errors happen while the program is running. They're worse because they can
crash a program that's already in use.

