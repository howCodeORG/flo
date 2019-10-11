package main

import (
	"flo/eval"
	"flo/lexer"
	"flo/parser"
	"flo/repl"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		// REPL
		// Initialise environment
		//eval.Init()
		repl.Start(os.Stdin, os.Stdout)
	} else {

		// Initialise environment
		eval.Init()
		file, err := os.Open(args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		// fmt.Print(string(b))

		l := lexer.New(string(b))
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
		p := parser.New(l)
		program := p.ParseProgram()
		// fmt.Println(program)
		if len(p.Errors()) == 0 {
			eval.Eval(program)
			if eval.Error != "" {
				fmt.Println(eval.Error + "\r\n")
				eval.Error = ""
			}
		} else {
			fmt.Println("-- SyntaxError(s) -- ")
			for _, error := range p.Errors() {
				fmt.Println("   " + error)
			}
		}

	}
}
