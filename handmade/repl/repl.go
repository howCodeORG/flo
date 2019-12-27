package repl

import (
	"flo/eval"
	"flo/lexer"
	"flo/parser"
	"fmt"
	"io"
	"os"
	"runtime/debug"

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

	// Initialise environment
	eval.Init()

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
		l := lexer.New(input)
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
		p := parser.New(l)
		program := p.ParseProgram()
		if program.String() != "" && len(p.Errors()) == 0 {
			// fmt.Println(program)
		}
		if len(p.Errors()) == 0 {
			p := eval.Eval(program)
			if p != nil {
				// fmt.Println(p)
			}
			if eval.Error != "" {
				fmt.Println(eval.Error)
				eval.Error = ""
			}
		} else {
			for _, error := range p.Errors() {
				fmt.Println(error)
			}
		}

	}

	// term, err := terminal.NewWithStdInOut()
	// if err != nil {
	// 	panic(err)
	// }
	// defer term.ReleaseFromStdInOut() // defer this
	// fmt.Println("-- Welcome to Flo (0.0.1) --\n-- Ctrl-C to exit --")
	// term.SetPrompt(">> ")
	// line, err := term.ReadLine()
	// for {
	// 	if err == io.EOF {
	// 		term.Write([]byte(line))
	// 		fmt.Println()
	// 		return
	// 	}
	// 	if err != nil && strings.Contains(err.Error(), "control-c break") {
	// 		term.Write([]byte(line))
	// 		fmt.Println()
	// 		return
	// 	} else {
	// 		l := lexer.New(line)
	// 		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 		// 	fmt.Printf("%+v\n", tok)
	// 		// }
	// 		p := parser.New(l)
	// 		program := p.ParseProgram()
	// 		// fmt.Println(program.GetValue())
	// 		if len(p.Errors()) == 0 {
	// 			eval.Eval(program)
	// 			if eval.Error != "" {
	// 				term.Write([]byte(eval.Error + "\r\n"))
	// 				eval.Error = ""
	// 			}
	// 		} else {
	// 			for _, error := range p.Errors() {
	// 				fmt.Println(error)
	// 			}
	// 		}
	// 		line, err = term.ReadLine()
	// 	}
	// }

}
