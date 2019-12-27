package compiler

import (
	"flo/parser"
	"flo/vm"
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func doEval(p *parser.FloParser, visitor *FloVisitor) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// Finally parse the expression
	antlr.ParseTreeVisitor.Visit(visitor, p.Start())

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

	output := [][]byte{
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.BINARY_ADD, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.BINARY_SUB, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.BINARY_MUL, 0,
			vm.LOAD_CONST, 2,
			vm.BINARY_SUB, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.STORE_NAME, 0,
			vm.LOAD_CONST, 1,
			vm.STORE_NAME, 1,
			vm.LOAD_NAME, 0,
			vm.LOAD_NAME, 1,
		},
		{
			vm.LOAD_CONST, 0,
			vm.STORE_NAME, 0,
			vm.LOAD_CONST, 1,
			vm.STORE_NAME, 1,
			vm.LOAD_NAME, 0,
			vm.LOAD_NAME, 1,
			vm.BINARY_ADD, 0,
			vm.STORE_NAME, 2,
			vm.LOAD_NAME, 2,
			vm.LOAD_CONST, 2,
			vm.BINARY_DIV, 0,
			vm.STORE_NAME, 3,
			vm.LOAD_NAME, 0,
			vm.LOAD_NAME, 1,
			vm.LOAD_NAME, 2,
			vm.LOAD_NAME, 3,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.BINARY_MUL, 0,
		},
		{
			vm.LOAD_CONST, 0,
		},
		{
			vm.LOAD_CONST, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.LOGICAL_OR, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 0,
			vm.LOGICAL_OR, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.LOGICAL_AND, 0,
		},
		{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 0,
			vm.LOGICAL_AND, 0,
		},
	}

	for i := range output {

		var visitor FloVisitor
		flovm := vm.VM{}
		flovm.Init()
		visitor.Init()

		// Setup the input
		is := antlr.NewInputStream(input[i])

		// Create the Lexer
		lexer := parser.NewFloLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewFloParser(stream)

		doEval(p, &visitor)

		for j := range output[i] {

			got := visitor.Object.Instructions[j]
			expected := output[i][j]
			if expected != got {
				t.Errorf("Test %d failed - Got %d, expected %d\n", i, got, expected)
			}

		}

	}

}
