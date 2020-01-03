The Flo Programming Language
============================
Flo stands for "Francis' Language written in Go'.

It's a project I've been working on to learn more about programming language design and compilers. Flo uses the ANTLR parser generator library to generate its lexer and parser.

Flo is compiled into bytecode running on a custom virtual machine inspired by the Python virtual machine, so Flo bytecode looks quite similar to Python bytecode.

You can compile Flo by typing:
`go build`

You'll most likely need to install some of Flo's dependencies namely

ANTLR

`go get github.com/antlr/antlr4/runtime/Go/antlr`

and go-prompt

`go get github.com/c-bata/go-prompt`

If there are any dependencies I've forgotten about, go will tell you when you run `go build`.

After you've built Flo, try it out!

`go build` will produce an executable called `flo` on Mac and Linux and `flo.exe` on Windows.

If you run the executable without passing any parameters you'll see the Flo REPL, it's an interactive prompt you can you type Flo code in your Terminal.
`./flo` will launch the REPL.

If you provide a file name like `./flo test.flo` Flo will load the file you provided and will execute it, just like when you run `python test.py` to run Python programs.

### Example programs
Look in the examples folder if you want to see the example code, I've copied and pasted some here.
#### Flo data types
Taken from `examples/data_types.flo`
```
// Flo is dynamically typed
x  =  5  // x is an int
x  =  true  // x is now a bool

// Lists can store any data type, including functions or other lists
mylist  = [1, 2, 3, 'hello world', true, false, ['nested list', 1,2,3]]
print mylist // This will print my list
f  =  func() {
print  'f is an anonymous function that was assigned to a variable'
}
f() // This will call f because functions are a first-class type
```
This will output:
```
[1, 2, 3, 'hello world', true, false, ['nested list', 1, 2, 3]]
'f is an anonymous function that was assigned to a variable'
```
#### Fibonacci
Taken from `examples/fib.flo`
```
// Recursive fibonacci function in Flo
func  fib(n) {
    if n <  2 {
        return n
    } else {
    return  fib(n-1) +  fib(n-2)
    }
}
print fib(10)
```
This will output
```
55
```

#### Flo bytecode
Here's an example of the Flo bytecode produced by the compiler for the Fibonacci function. These are the actual instructions that run inside the virtual machine when you execute your program.
```
Instructions (disassembled):
0 SETUP_FUNCTION 0
2 LOAD_CONST 'n' (0)
4 SETUP_PARAMS 1
6 SETUP_BODY 0
8 PUSH_BLOCK 0
10 LOAD_NAME 'n' (1)
12 LOAD_CONST 2 (1)
14 COMPARE 0
16 POP_JUMP_IF_FALSE 8
18 LOAD_NAME 'n' (1)
20 RETURN 0
22 JUMP 26
24 LOAD_NAME 'n' (1)
26 LOAD_CONST 1 (2)
28 BINARY_SUB 0
30 LOAD_NAME 'fib' (0)
32 CALL_FUNCTION 1
34 LOAD_NAME 'n' (1)
36 LOAD_CONST 2 (1)
38 BINARY_SUB 0
40 LOAD_NAME 'fib' (0)
42 CALL_FUNCTION 1
44 BINARY_ADD 0
46 RETURN 0
48 POP_BLOCK 0
50 MAKE_FUNCTION 0
52 STORE_NAME 'fib' (0)
Instructions (bytecode):
8 0 0 0 9 1 10 0 4 0 1 1 0 1 30 0 31 8 1 1 13 0 35 26 1 1 0 
2 15 0 1 0 12 1 1 1 0 1 15 0 1 0 12 1 14 0 13 0 5 0 11 0 2 0
```
