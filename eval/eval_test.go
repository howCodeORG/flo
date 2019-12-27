package eval

import (
	"flo/parser"
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func doEval(p *parser.FloParser, listener *FloListener) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

}

func TestExpressions(t *testing.T) {

	input := []string{
		"1+2",
		"5-1",
		"100 * 2 - 4",
		"a = 1; b = 2; a; b;",
		"a = 10; b = 5; c = a + b; d = c / 3; a;b; c; d;",
		"10.5 * 2",
		"true",
		"false",
		"true or false",
		"false or false",
		"true and false",
		"true and true",
	}

	output := [][]FloObject{
		{FloInteger(3)},
		{FloInteger(4)},
		{FloInteger(196)},
		{FloInteger(1), FloInteger(2)},
		{FloInteger(10), FloInteger(5), FloInteger(15), FloInteger(5)},
		{FloFloat(21)},
		{FloBool(true)},
		{FloBool(false)},
		{FloBool(true)},
		{FloBool(false)},
		{FloBool(false)},
		{FloBool(true)},
	}

	for i := range output {

		var listener FloListener
		listener.Init()

		// Setup the input
		is := antlr.NewInputStream(input[i])

		// Create the Lexer
		lexer := parser.NewFloLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewFloParser(stream)

		doEval(p, &listener)

		for j := range output[i] {

			got := listener.Out()
			expected := output[i][len(output[i])-1-j]
			if expected != got {
				t.Errorf("Got %s, expected %s\n", got, expected)
			}

		}

	}

}
