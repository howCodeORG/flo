package lexer

import (
	"flo/token"
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `100 + 200; test;`

	tests := []struct {
		expectedType  string
		expectedValue string
	}{
		{token.INT, "100"},
		{token.PLUS, "+"},
		{token.INT, "200"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "test"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, test := range tests {

		tok := l.NextToken()

		if test.expectedType != tok.Type {
			t.Fatalf("tests[%d] - Incorrect token type - expected %q, got %q", i, test.expectedType, tok.Type)
		}

		if test.expectedValue != tok.Value {
			t.Fatalf("tests[%d] - Incorrect token value - expected %q, got %q", i, test.expectedValue, tok.Value)
		}

		fmt.Println(tok)

	}
}
