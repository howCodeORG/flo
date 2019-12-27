package main

import (
	"flo/vm"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"flo/eval"
	"flo/parser"

	"github.com/c-bata/go-prompt"
)

type ErrorListener struct {
}

func NewErrorListener() *ErrorListener {
	return new(ErrorListener)
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic("SyntaxError: invalid syntax, " + "line " + strconv.Itoa(line) + ", col " + strconv.Itoa(column) + " " + msg)
	// fmt.Println(e.GetInputStream().Index())
	// fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func (d *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (d *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (d *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		// {Text: "users", Description: "Store the username and age"},
		// {Text: "articles", Description: "Store the article text posted by user"},
		// {Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

type Exit int

func exit(_ *prompt.Buffer) {
	panic(Exit(0))
}

func handleExit() {
	switch v := recover().(type) {
	case nil:
		return
	case Exit:
		os.Exit(int(v))
	default:
		fmt.Println(v)
		// fmt.Println(string(debug.Stack()))
	}
}

func doEval(p *parser.FloParser, visitor *eval.FloVisitor) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("-- traceback --")
			for i, _ := range visitor.Stack {
				fmt.Println(visitor.Stack[len(visitor.Stack)-i-1])
			}
			fmt.Println(r)
			// fmt.Println(string(debug.Stack()))
		}
	}()

	// Finally parse the expression
	// antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())
	antlr.ParseTreeVisitor.Visit(visitor, p.Start())
	// out := listener.Out()
	// if out != nil {
	// 	fmt.Println(out)
	// }
}

func Start(in io.Reader, out io.Writer) {

	defer handleExit()

	// Initialise history slice
	history := []string{}

	fmt.Println("-- Welcome to Flo (0.0.1) --\n-- Ctrl-C to exit --")

	// var listener eval.FloListener
	var visitor eval.FloVisitor
	visitor.Init()
	flovm := vm.VM{}
	flovm.Init()

	for {

		input := prompt.Input(">> ", completer,
			prompt.OptionHistory(history),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn:  exit,
			}))

		history = append(history, input)

		if input == "" {
			continue
		}

		input += "\n"

		// Setup the input
		is := antlr.NewInputStream(input)

		// Create the Lexer
		lexer := parser.NewFloLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewFloParser(stream)

		p.RemoveErrorListeners()

		e := NewErrorListener()

		p.AddErrorListener(e)

		doEval(p, &visitor)

		f := vm.FloCallable{
			Args:   ([]vm.FloObject{}),
			Object: visitor.Object,
			Name:   vm.FloString("main"),
		}
		flovm.Run(f, vm.FloList([]vm.FloObject{}))
		visitor.Object.Instructions = nil

	}

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	args := os.Args

	if len(args) < 2 {
		// REPL
		Start(os.Stdin, os.Stdout)
	} else {

		var visitor eval.FloVisitor
		flovm := vm.VM{}
		flovm.Init()
		visitor.Init()

		file, err := os.Open(args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)

		input := string(b)

		if input == "" {
			return
		}

		// Setup the input
		is := antlr.NewInputStream(input)

		// Create the Lexer
		lexer := parser.NewFloLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		// Create the Parser
		p := parser.NewFloParser(stream)

		p.RemoveErrorListeners()

		e := NewErrorListener()

		p.AddErrorListener(e)

		doEval(p, &visitor)
		f := vm.FloCallable{
			Args:   ([]vm.FloObject{}),
			Object: visitor.Object,
			Name:   vm.FloString("main"),
		}
		flovm.Run(f, vm.FloList([]vm.FloObject{}))

	}
}
