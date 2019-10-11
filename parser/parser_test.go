package parser

//
// import (
// 	"lang/lexer"
// 	"testing"
// )
//
// func TestParser(t *testing.T) {
//
// 	inputs := []string{
// 		"100",
// 		"1+2",
// 		"1+2+3+4+5",
// 		"1+2*3**4",
// 		"-1",
// 		"--2",
// 		"1*2*4*5*6",
// 		"100+200; 500+300",
// 		"1;2;3;",
// 	}
//
// 	tests := []struct {
// 		expectedTree string
// 	}{
// 		{"[RootNode] { [StatementNode] { [IntegerNode] {Value: 100}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [BinaryOpNode] {Operator: +, Left: [IntegerNode] {Value: 1}, Right: [IntegerNode] {Value: 2}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [BinaryOpNode] {Operator: +, Left: [BinaryOpNode] {Operator: +, Left: [BinaryOpNode] {Operator: +, Left: [BinaryOpNode] {Operator: +, Left: [IntegerNode] {Value: 1}, Right: [IntegerNode] {Value: 2}}, Right: [IntegerNode] {Value: 3}}, Right: [IntegerNode] {Value: 4}}, Right: [IntegerNode] {Value: 5}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [BinaryOpNode] {Operator: +, Left: [IntegerNode] {Value: 1}, Right: [BinaryOpNode] {Operator: *, Left: [IntegerNode] {Value: 2}, Right: [BinaryOpNode] {Operator: **, Left: [IntegerNode] {Value: 3}, Right: [IntegerNode] {Value: 4}}}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [UnaryOpNode] {Operator: -, Value: [IntegerNode] {Value: 1}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [UnaryOpNode] {Operator: -, Value: [UnaryOpNode] {Operator: -, Value: [IntegerNode] {Value: 2}}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [BinaryOpNode] {Operator: *, Left: [BinaryOpNode] {Operator: *, Left: [BinaryOpNode] {Operator: *, Left: [BinaryOpNode] {Operator: *, Left: [IntegerNode] {Value: 1}, Right: [IntegerNode] {Value: 2}}, Right: [IntegerNode] {Value: 4}}, Right: [IntegerNode] {Value: 5}}, Right: [IntegerNode] {Value: 6}}, Next: nil } }"},
// 		{"[RootNode] { [StatementNode] { [BinaryOpNode] {Operator: +, Left: [IntegerNode] {Value: 100}, Right: [IntegerNode] {Value: 200}}, Next: [StatementNode] { [BinaryOpNode] {Operator: +, Left: [IntegerNode] {Value: 500}, Right: [IntegerNode] {Value: 300}}, Next: nil } } }"},
// 		{"[RootNode] { [StatementNode] { [IntegerNode] {Value: 1}, Next: [StatementNode] { [IntegerNode] {Value: 2}, Next: [StatementNode] { [IntegerNode] {Value: 3}, Next: nil } } } }"},
// 	}
//
// 	for i, test := range tests {
// 		l := lexer.New(inputs[i])
// 		p := New(l)
// 		output := p.ParseProgram().GetValue()
//
// 		if test.expectedTree != output {
// 			t.Fatalf("tests[%d] - Invalid parse tree - expected %q, got %q", i, test.expectedTree, output)
// 		}
// 	}
//
// }
