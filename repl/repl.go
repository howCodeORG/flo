package repl

import (
	"flo/parser"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"
)

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
		fmt.Println(string(debug.Stack()))
	}
}

func Start(in io.Reader, out io.Writer) {

	defer handleExit()

	// Initialise history slice
	history := []string{}

	fmt.Println("-- Welcome to Flo (0.0.1) --\n-- Ctrl-C to exit --")

	for {
		input := prompt.Input(">> ", completer,
			prompt.OptionHistory(history),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn:  exit,
			}))

		history = append(history, input)

		// Setup the input
		is := antlr.NewInputStream(input)

		// Create the Lexer
		lexer := parser.NewFloLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewFloParser(stream)

		var listener floListener
		// Finally parse the expression
		antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

		fmt.Println(listener.pop())

		// l := lexer.New(input)
		// l.Debug()
		// p := parser.New(l)
		// tree := p.Parse()

		// r.Run(tree)

	}

}
