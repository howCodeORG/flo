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
y  =  true  // x is not a bool

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
// Recursive fibonacci function in Flo
```
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
