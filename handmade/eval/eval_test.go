package eval

import (
	"testing"
)

func TestEvaluation(t *testing.T) {

	// inputs := []string{
	// 	"100",
	// 	"1+2",
	// 	"1+2+3+4+5",
	// 	"1+2*3**4",
	// 	"-1",
	// 	"--2",
	// 	"1*2*4*5*6",
	// 	"a = 10; a + 20",
	// }
	//
	// tests := []string{
	// 	"100",
	// 	"3",
	// 	"15",
	// 	"163",
	// 	"-1",
	// 	"2",
	// 	"240",
	// 	"",
	// }
	//
	// for i, test := range tests {
	// 	l := lexer.New(inputs[i])
	// 	program := parser.New(l).ParseProgram()
	// 	// fmt.Println(program.GetValue())
	// 	// Create environment to hold state
	// 	Init()
	// 	response := Eval(program)
	// 	output := ""
	// 	if response != nil {
	// 		output = response.Inspect()
	// 	}
	// 	if test != output {
	// 		t.Fatalf("tests[%d] - Code executed incorrectly - expected %q, got %q", i, test, output)
	// 	}
	// }

}
