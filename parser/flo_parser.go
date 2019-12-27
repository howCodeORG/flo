// Generated from Flo.g4 by ANTLR 4.7.

package parser // Flo

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 220,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 6, 3, 43, 10, 3, 13, 3, 14, 3,
	44, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 54, 10, 4, 3, 5, 3,
	5, 5, 5, 58, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7,
	7, 68, 10, 7, 12, 7, 14, 7, 71, 11, 7, 3, 7, 3, 7, 3, 7, 5, 7, 76, 10,
	7, 3, 7, 5, 7, 79, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 88, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 93, 10, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 7, 10, 100, 10, 10, 12, 10, 14, 10, 103, 11, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 112, 10, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 5, 12, 119, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 5, 14, 128, 10, 14, 3, 15, 3, 15, 5, 15, 132, 10, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 149, 10, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 5, 16, 158, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 193, 10,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 201, 10, 16, 12, 16,
	14, 16, 204, 11, 16, 3, 17, 3, 17, 3, 17, 7, 17, 209, 10, 17, 12, 17, 14,
	17, 212, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 218, 10, 18, 3, 18,
	2, 3, 30, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 2, 7, 3, 2, 35, 36, 3, 2, 33, 34, 5, 2, 25, 25, 30, 30, 32, 32, 3,
	2, 26, 27, 3, 2, 18, 23, 2, 248, 2, 36, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2,
	6, 53, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 63, 3, 2,
	2, 2, 14, 80, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20, 104,
	3, 2, 2, 2, 22, 115, 3, 2, 2, 2, 24, 122, 3, 2, 2, 2, 26, 125, 3, 2, 2,
	2, 28, 129, 3, 2, 2, 2, 30, 157, 3, 2, 2, 2, 32, 205, 3, 2, 2, 2, 34, 217,
	3, 2, 2, 2, 36, 37, 5, 4, 3, 2, 37, 38, 7, 2, 2, 3, 38, 3, 3, 2, 2, 2,
	39, 40, 5, 6, 4, 2, 40, 41, 5, 34, 18, 2, 41, 43, 3, 2, 2, 2, 42, 39, 3,
	2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45,
	5, 3, 2, 2, 2, 46, 54, 5, 8, 5, 2, 47, 54, 5, 12, 7, 2, 48, 54, 5, 14,
	8, 2, 49, 54, 5, 16, 9, 2, 50, 54, 5, 24, 13, 2, 51, 54, 5, 26, 14, 2,
	52, 54, 5, 28, 15, 2, 53, 46, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3,
	2, 2, 2, 53, 49, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53,
	52, 3, 2, 2, 2, 54, 7, 3, 2, 2, 2, 55, 58, 5, 30, 16, 2, 56, 58, 5, 10,
	6, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 9, 3, 2, 2, 2, 59, 60,
	7, 51, 2, 2, 60, 61, 7, 3, 2, 2, 61, 62, 5, 30, 16, 2, 62, 11, 3, 2, 2,
	2, 63, 64, 7, 10, 2, 2, 64, 69, 5, 22, 12, 2, 65, 66, 7, 13, 2, 2, 66,
	68, 5, 22, 12, 2, 67, 65, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2,
	2, 2, 69, 70, 3, 2, 2, 2, 70, 78, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73,
	7, 12, 2, 2, 73, 75, 7, 41, 2, 2, 74, 76, 5, 4, 3, 2, 75, 74, 3, 2, 2,
	2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 7, 42, 2, 2, 78, 72,
	3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 13, 3, 2, 2, 2, 80, 81, 7, 11, 2, 2,
	81, 82, 5, 20, 11, 2, 82, 15, 3, 2, 2, 2, 83, 84, 7, 16, 2, 2, 84, 85,
	7, 51, 2, 2, 85, 87, 7, 45, 2, 2, 86, 88, 5, 18, 10, 2, 87, 86, 3, 2, 2,
	2, 87, 88, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 46, 2, 2, 90, 92,
	7, 41, 2, 2, 91, 93, 5, 4, 3, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2,
	93, 94, 3, 2, 2, 2, 94, 95, 7, 42, 2, 2, 95, 17, 3, 2, 2, 2, 96, 101, 7,
	51, 2, 2, 97, 98, 7, 4, 2, 2, 98, 100, 7, 51, 2, 2, 99, 97, 3, 2, 2, 2,
	100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 19,
	3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 105, 5, 8, 5, 2, 105, 106, 7, 47,
	2, 2, 106, 107, 5, 30, 16, 2, 107, 108, 7, 47, 2, 2, 108, 109, 5, 8, 5,
	2, 109, 111, 7, 41, 2, 2, 110, 112, 5, 4, 3, 2, 111, 110, 3, 2, 2, 2, 111,
	112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 7, 42, 2, 2, 114, 21,
	3, 2, 2, 2, 115, 116, 5, 30, 16, 2, 116, 118, 7, 41, 2, 2, 117, 119, 5,
	4, 3, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 3, 2, 2,
	2, 120, 121, 7, 42, 2, 2, 121, 23, 3, 2, 2, 2, 122, 123, 7, 9, 2, 2, 123,
	124, 5, 30, 16, 2, 124, 25, 3, 2, 2, 2, 125, 127, 7, 17, 2, 2, 126, 128,
	5, 30, 16, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 27, 3, 2,
	2, 2, 129, 131, 7, 41, 2, 2, 130, 132, 5, 4, 3, 2, 131, 130, 3, 2, 2, 2,
	131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 7, 42, 2, 2, 134,
	29, 3, 2, 2, 2, 135, 136, 8, 16, 1, 2, 136, 137, 7, 45, 2, 2, 137, 138,
	5, 30, 16, 2, 138, 139, 7, 46, 2, 2, 139, 158, 3, 2, 2, 2, 140, 141, 9,
	2, 2, 2, 141, 158, 5, 30, 16, 22, 142, 143, 7, 51, 2, 2, 143, 158, 9, 3,
	2, 2, 144, 145, 7, 29, 2, 2, 145, 158, 5, 30, 16, 12, 146, 148, 7, 43,
	2, 2, 147, 149, 5, 32, 17, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2,
	2, 149, 150, 3, 2, 2, 2, 150, 158, 7, 44, 2, 2, 151, 158, 7, 48, 2, 2,
	152, 158, 7, 50, 2, 2, 153, 158, 7, 51, 2, 2, 154, 158, 7, 7, 2, 2, 155,
	158, 7, 8, 2, 2, 156, 158, 7, 49, 2, 2, 157, 135, 3, 2, 2, 2, 157, 140,
	3, 2, 2, 2, 157, 142, 3, 2, 2, 2, 157, 144, 3, 2, 2, 2, 157, 146, 3, 2,
	2, 2, 157, 151, 3, 2, 2, 2, 157, 152, 3, 2, 2, 2, 157, 153, 3, 2, 2, 2,
	157, 154, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 156, 3, 2, 2, 2, 158,
	202, 3, 2, 2, 2, 159, 160, 12, 23, 2, 2, 160, 161, 7, 31, 2, 2, 161, 201,
	5, 30, 16, 24, 162, 163, 12, 19, 2, 2, 163, 164, 9, 4, 2, 2, 164, 201,
	5, 30, 16, 20, 165, 166, 12, 18, 2, 2, 166, 167, 9, 2, 2, 2, 167, 201,
	5, 30, 16, 19, 168, 169, 12, 17, 2, 2, 169, 170, 9, 5, 2, 2, 170, 201,
	5, 30, 16, 18, 171, 172, 12, 16, 2, 2, 172, 173, 7, 5, 2, 2, 173, 201,
	5, 30, 16, 17, 174, 175, 12, 15, 2, 2, 175, 176, 7, 6, 2, 2, 176, 201,
	5, 30, 16, 16, 177, 178, 12, 14, 2, 2, 178, 179, 7, 24, 2, 2, 179, 201,
	5, 30, 16, 15, 180, 181, 12, 13, 2, 2, 181, 182, 9, 6, 2, 2, 182, 201,
	5, 30, 16, 14, 183, 184, 12, 11, 2, 2, 184, 185, 7, 15, 2, 2, 185, 201,
	5, 30, 16, 12, 186, 187, 12, 10, 2, 2, 187, 188, 7, 14, 2, 2, 188, 201,
	5, 30, 16, 11, 189, 190, 12, 24, 2, 2, 190, 192, 7, 45, 2, 2, 191, 193,
	5, 32, 17, 2, 192, 191, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3,
	2, 2, 2, 194, 201, 7, 46, 2, 2, 195, 196, 12, 20, 2, 2, 196, 197, 7, 43,
	2, 2, 197, 198, 5, 30, 16, 2, 198, 199, 7, 44, 2, 2, 199, 201, 3, 2, 2,
	2, 200, 159, 3, 2, 2, 2, 200, 162, 3, 2, 2, 2, 200, 165, 3, 2, 2, 2, 200,
	168, 3, 2, 2, 2, 200, 171, 3, 2, 2, 2, 200, 174, 3, 2, 2, 2, 200, 177,
	3, 2, 2, 2, 200, 180, 3, 2, 2, 2, 200, 183, 3, 2, 2, 2, 200, 186, 3, 2,
	2, 2, 200, 189, 3, 2, 2, 2, 200, 195, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2,
	202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 31, 3, 2, 2, 2, 204, 202,
	3, 2, 2, 2, 205, 210, 5, 30, 16, 2, 206, 207, 7, 4, 2, 2, 207, 209, 5,
	30, 16, 2, 208, 206, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2,
	2, 2, 210, 211, 3, 2, 2, 2, 211, 33, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2,
	213, 218, 7, 47, 2, 2, 214, 218, 7, 2, 2, 3, 215, 218, 6, 18, 14, 2, 216,
	218, 6, 18, 15, 2, 217, 213, 3, 2, 2, 2, 217, 214, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 35, 3, 2, 2, 2, 22, 44, 53, 57,
	69, 75, 78, 87, 92, 101, 111, 118, 127, 131, 148, 157, 192, 200, 202, 210,
	217,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "','", "'&'", "'^'", "'true'", "'false'", "'print'", "'if'",
	"'for'", "'else'", "'elif'", "'or'", "'and'", "'func'", "'return'", "'=='",
	"'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'%'", "'<<'", "'>>'", "'&^'",
	"'not'", "'*'", "'**'", "'/'", "'++'", "'--'", "'+'", "'-'", "", "", "",
	"", "'{'", "'}'", "'['", "']'", "'('", "')'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "TRUE", "FALSE", "PRINT", "IF", "FOR", "ELSE", "ELIF",
	"LOGICAL_OR", "LOGICAL_AND", "FUNC", "RETURN", "EQUALS", "NOT_EQUALS",
	"LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "MOD",
	"LSHIFT", "RSHIFT", "BIT_CLEAR", "NOT", "MULTIPLY", "POWER", "DIVIDE",
	"PLUS_PLUS", "MINUS_MINUS", "ADDITION", "SUBTRACTION", "WS", "COMMENT",
	"TERMINATOR", "LINE_COMMENT", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET",
	"L_PAREN", "R_PAREN", "SEMI", "INTEGER", "CHARS", "FLOATING_POINT", "IDENTIFIER",
}

var ruleNames = []string{
	"start", "multi_stmts", "stmt", "simple_stmt", "assign_stmt", "if_stmt",
	"for_stmt", "func_decl", "parameters", "for_clause_block", "condition_block",
	"print_stmt", "return_stmt", "block", "expression", "expression_list",
	"eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FloParser struct {
	FloParserBase
}

func NewFloParser(input antlr.TokenStream) *FloParser {
	this := new(FloParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Flo.g4"

	return this
}

// FloParser tokens.
const (
	FloParserEOF               = antlr.TokenEOF
	FloParserT__0              = 1
	FloParserT__1              = 2
	FloParserT__2              = 3
	FloParserT__3              = 4
	FloParserTRUE              = 5
	FloParserFALSE             = 6
	FloParserPRINT             = 7
	FloParserIF                = 8
	FloParserFOR               = 9
	FloParserELSE              = 10
	FloParserELIF              = 11
	FloParserLOGICAL_OR        = 12
	FloParserLOGICAL_AND       = 13
	FloParserFUNC              = 14
	FloParserRETURN            = 15
	FloParserEQUALS            = 16
	FloParserNOT_EQUALS        = 17
	FloParserLESS              = 18
	FloParserLESS_OR_EQUALS    = 19
	FloParserGREATER           = 20
	FloParserGREATER_OR_EQUALS = 21
	FloParserOR                = 22
	FloParserMOD               = 23
	FloParserLSHIFT            = 24
	FloParserRSHIFT            = 25
	FloParserBIT_CLEAR         = 26
	FloParserNOT               = 27
	FloParserMULTIPLY          = 28
	FloParserPOWER             = 29
	FloParserDIVIDE            = 30
	FloParserPLUS_PLUS         = 31
	FloParserMINUS_MINUS       = 32
	FloParserADDITION          = 33
	FloParserSUBTRACTION       = 34
	FloParserWS                = 35
	FloParserCOMMENT           = 36
	FloParserTERMINATOR        = 37
	FloParserLINE_COMMENT      = 38
	FloParserL_CURLY           = 39
	FloParserR_CURLY           = 40
	FloParserL_BRACKET         = 41
	FloParserR_BRACKET         = 42
	FloParserL_PAREN           = 43
	FloParserR_PAREN           = 44
	FloParserSEMI              = 45
	FloParserINTEGER           = 46
	FloParserCHARS             = 47
	FloParserFLOATING_POINT    = 48
	FloParserIDENTIFIER        = 49
)

// FloParser rules.
const (
	FloParserRULE_start            = 0
	FloParserRULE_multi_stmts      = 1
	FloParserRULE_stmt             = 2
	FloParserRULE_simple_stmt      = 3
	FloParserRULE_assign_stmt      = 4
	FloParserRULE_if_stmt          = 5
	FloParserRULE_for_stmt         = 6
	FloParserRULE_func_decl        = 7
	FloParserRULE_parameters       = 8
	FloParserRULE_for_clause_block = 9
	FloParserRULE_condition_block  = 10
	FloParserRULE_print_stmt       = 11
	FloParserRULE_return_stmt      = 12
	FloParserRULE_block            = 13
	FloParserRULE_expression       = 14
	FloParserRULE_expression_list  = 15
	FloParserRULE_eos              = 16
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(FloParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FloParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Multi_stmts()
	}
	{
		p.SetState(35)
		p.Match(FloParserEOF)
	}

	return localctx
}

// IMulti_stmtsContext is an interface to support dynamic dispatch.
type IMulti_stmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulti_stmtsContext differentiates from other interfaces.
	IsMulti_stmtsContext()
}

type Multi_stmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulti_stmtsContext() *Multi_stmtsContext {
	var p = new(Multi_stmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_multi_stmts
	return p
}

func (*Multi_stmtsContext) IsMulti_stmtsContext() {}

func NewMulti_stmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multi_stmtsContext {
	var p = new(Multi_stmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_multi_stmts

	return p
}

func (s *Multi_stmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *Multi_stmtsContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *Multi_stmtsContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Multi_stmtsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *Multi_stmtsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *Multi_stmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multi_stmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multi_stmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterMulti_stmts(s)
	}
}

func (s *Multi_stmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitMulti_stmts(s)
	}
}

func (s *Multi_stmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitMulti_stmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Multi_stmts() (localctx IMulti_stmtsContext) {
	localctx = NewMulti_stmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FloParserRULE_multi_stmts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(37)
			p.Stmt()
		}
		{
			p.SetState(38)
			p.Eos()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Simple_stmt() ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Func_decl() IFunc_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_declContext)
}

func (s *StmtContext) Print_stmt() IPrint_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint_stmtContext)
}

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FloParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FloParserTRUE, FloParserFALSE, FloParserNOT, FloParserADDITION, FloParserSUBTRACTION, FloParserL_BRACKET, FloParserL_PAREN, FloParserINTEGER, FloParserCHARS, FloParserFLOATING_POINT, FloParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Simple_stmt()
		}

	case FloParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.If_stmt()
		}

	case FloParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.For_stmt()
		}

	case FloParserFUNC:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Func_decl()
		}

	case FloParserPRINT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Print_stmt()
		}

	case FloParserRETURN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Return_stmt()
		}

	case FloParserL_CURLY:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(50)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimple_stmtContext is an interface to support dynamic dispatch.
type ISimple_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_stmtContext differentiates from other interfaces.
	IsSimple_stmtContext()
}

type Simple_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_stmtContext() *Simple_stmtContext {
	var p = new(Simple_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_simple_stmt
	return p
}

func (*Simple_stmtContext) IsSimple_stmtContext() {}

func NewSimple_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_stmtContext {
	var p = new(Simple_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_simple_stmt

	return p
}

func (s *Simple_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Simple_stmtContext) Assign_stmt() IAssign_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *Simple_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterSimple_stmt(s)
	}
}

func (s *Simple_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitSimple_stmt(s)
	}
}

func (s *Simple_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitSimple_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Simple_stmt() (localctx ISimple_stmtContext) {
	localctx = NewSimple_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FloParserRULE_simple_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Assign_stmt()
		}

	}

	return localctx
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_assign_stmt
	return p
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FloParserIDENTIFIER, 0)
}

func (s *Assign_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitAssign_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FloParserRULE_assign_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(FloParserIDENTIFIER)
	}
	{
		p.SetState(58)
		p.Match(FloParserT__0)
	}
	{
		p.SetState(59)
		p.expression(0)
	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(FloParserIF, 0)
}

func (s *If_stmtContext) AllCondition_block() []ICondition_blockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem())
	var tst = make([]ICondition_blockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondition_blockContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Condition_block(i int) ICondition_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondition_blockContext)
}

func (s *If_stmtContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(FloParserELIF)
}

func (s *If_stmtContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(FloParserELIF, i)
}

func (s *If_stmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FloParserELSE, 0)
}

func (s *If_stmtContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FloParserRULE_if_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(FloParserIF)
	}
	{
		p.SetState(62)
		p.Condition_block()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(63)
				p.Match(FloParserELIF)
			}
			{
				p.SetState(64)
				p.Condition_block()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(70)
			p.Match(FloParserELSE)
		}
		{
			p.SetState(71)
			p.Match(FloParserL_CURLY)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
			{
				p.SetState(72)
				p.Multi_stmts()
			}

		}
		{
			p.SetState(75)
			p.Match(FloParserR_CURLY)
		}

	}

	return localctx
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_for_stmt
	return p
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(FloParserFOR, 0)
}

func (s *For_stmtContext) For_clause_block() IFor_clause_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_clause_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_clause_blockContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (s *For_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitFor_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FloParserRULE_for_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(FloParserFOR)
	}
	{
		p.SetState(79)
		p.For_clause_block()
	}

	return localctx
}

// IFunc_declContext is an interface to support dynamic dispatch.
type IFunc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_declContext differentiates from other interfaces.
	IsFunc_declContext()
}

type Func_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declContext() *Func_declContext {
	var p = new(Func_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_func_decl
	return p
}

func (*Func_declContext) IsFunc_declContext() {}

func NewFunc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declContext {
	var p = new(Func_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_func_decl

	return p
}

func (s *Func_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FloParserFUNC, 0)
}

func (s *Func_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FloParserIDENTIFIER, 0)
}

func (s *Func_declContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *Func_declContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *Func_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterFunc_decl(s)
	}
}

func (s *Func_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitFunc_decl(s)
	}
}

func (s *Func_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitFunc_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Func_decl() (localctx IFunc_declContext) {
	localctx = NewFunc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FloParserRULE_func_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(FloParserFUNC)
	}
	{
		p.SetState(82)
		p.Match(FloParserIDENTIFIER)
	}
	{
		p.SetState(83)
		p.Match(FloParserL_PAREN)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FloParserIDENTIFIER {
		{
			p.SetState(84)
			p.Parameters()
		}

	}
	{
		p.SetState(87)
		p.Match(FloParserR_PAREN)
	}
	{
		p.SetState(88)
		p.Match(FloParserL_CURLY)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(89)
			p.Multi_stmts()
		}

	}
	{
		p.SetState(92)
		p.Match(FloParserR_CURLY)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FloParserIDENTIFIER)
}

func (s *ParametersContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FloParserIDENTIFIER, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FloParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(FloParserIDENTIFIER)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FloParserT__1 {
		{
			p.SetState(95)
			p.Match(FloParserT__1)
		}
		{
			p.SetState(96)
			p.Match(FloParserIDENTIFIER)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFor_clause_blockContext is an interface to support dynamic dispatch.
type IFor_clause_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_clause_blockContext differentiates from other interfaces.
	IsFor_clause_blockContext()
}

type For_clause_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_clause_blockContext() *For_clause_blockContext {
	var p = new(For_clause_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_for_clause_block
	return p
}

func (*For_clause_blockContext) IsFor_clause_blockContext() {}

func NewFor_clause_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_clause_blockContext {
	var p = new(For_clause_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_for_clause_block

	return p
}

func (s *For_clause_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *For_clause_blockContext) AllSimple_stmt() []ISimple_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem())
	var tst = make([]ISimple_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimple_stmtContext)
		}
	}

	return tst
}

func (s *For_clause_blockContext) Simple_stmt(i int) ISimple_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *For_clause_blockContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(FloParserSEMI)
}

func (s *For_clause_blockContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(FloParserSEMI, i)
}

func (s *For_clause_blockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_clause_blockContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *For_clause_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_clause_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_clause_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterFor_clause_block(s)
	}
}

func (s *For_clause_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitFor_clause_block(s)
	}
}

func (s *For_clause_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitFor_clause_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) For_clause_block() (localctx IFor_clause_blockContext) {
	localctx = NewFor_clause_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FloParserRULE_for_clause_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Simple_stmt()
	}
	{
		p.SetState(103)
		p.Match(FloParserSEMI)
	}
	{
		p.SetState(104)
		p.expression(0)
	}
	{
		p.SetState(105)
		p.Match(FloParserSEMI)
	}
	{
		p.SetState(106)
		p.Simple_stmt()
	}
	{
		p.SetState(107)
		p.Match(FloParserL_CURLY)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(108)
			p.Multi_stmts()
		}

	}
	{
		p.SetState(111)
		p.Match(FloParserR_CURLY)
	}

	return localctx
}

// ICondition_blockContext is an interface to support dynamic dispatch.
type ICondition_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondition_blockContext differentiates from other interfaces.
	IsCondition_blockContext()
}

type Condition_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_blockContext() *Condition_blockContext {
	var p = new(Condition_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_condition_block
	return p
}

func (*Condition_blockContext) IsCondition_blockContext() {}

func NewCondition_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_blockContext {
	var p = new(Condition_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_condition_block

	return p
}

func (s *Condition_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_blockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Condition_blockContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *Condition_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterCondition_block(s)
	}
}

func (s *Condition_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitCondition_block(s)
	}
}

func (s *Condition_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitCondition_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Condition_block() (localctx ICondition_blockContext) {
	localctx = NewCondition_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FloParserRULE_condition_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.expression(0)
	}
	{
		p.SetState(114)
		p.Match(FloParserL_CURLY)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(115)
			p.Multi_stmts()
		}

	}
	{
		p.SetState(118)
		p.Match(FloParserR_CURLY)
	}

	return localctx
}

// IPrint_stmtContext is an interface to support dynamic dispatch.
type IPrint_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrint_stmtContext differentiates from other interfaces.
	IsPrint_stmtContext()
}

type Print_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_stmtContext() *Print_stmtContext {
	var p = new(Print_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_print_stmt
	return p
}

func (*Print_stmtContext) IsPrint_stmtContext() {}

func NewPrint_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_stmtContext {
	var p = new(Print_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_print_stmt

	return p
}

func (s *Print_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_stmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(FloParserPRINT, 0)
}

func (s *Print_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Print_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterPrint_stmt(s)
	}
}

func (s *Print_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitPrint_stmt(s)
	}
}

func (s *Print_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitPrint_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Print_stmt() (localctx IPrint_stmtContext) {
	localctx = NewPrint_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FloParserRULE_print_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(FloParserPRINT)
	}
	{
		p.SetState(121)
		p.expression(0)
	}

	return localctx
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_return_stmt
	return p
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(FloParserRETURN, 0)
}

func (s *Return_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (s *Return_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitReturn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FloParserRULE_return_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(FloParserRETURN)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.expression(0)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Multi_stmts() IMulti_stmtsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulti_stmtsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulti_stmtsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FloParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(FloParserL_CURLY)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserPRINT)|(1<<FloParserIF)|(1<<FloParserFOR)|(1<<FloParserFUNC)|(1<<FloParserRETURN)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_CURLY-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(128)
			p.Multi_stmts()
		}

	}
	{
		p.SetState(131)
		p.Match(FloParserR_CURLY)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryIncDecContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewUnaryIncDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryIncDecContext {
	var p = new(UnaryIncDecContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryIncDecContext) GetOp() antlr.Token { return s.op }

func (s *UnaryIncDecContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryIncDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryIncDecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FloParserIDENTIFIER, 0)
}

func (s *UnaryIncDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterUnaryIncDec(s)
	}
}

func (s *UnaryIncDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitUnaryIncDec(s)
	}
}

func (s *UnaryIncDecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitUnaryIncDec(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReadIdentifierContext struct {
	*ExpressionContext
}

func NewReadIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadIdentifierContext {
	var p = new(ReadIdentifierContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ReadIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FloParserIDENTIFIER, 0)
}

func (s *ReadIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterReadIdentifier(s)
	}
}

func (s *ReadIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitReadIdentifier(s)
	}
}

func (s *ReadIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitReadIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	*ExpressionContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(FloParserLOGICAL_OR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitShiftContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewBitShiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftContext {
	var p = new(BitShiftContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BitShiftContext) GetOp() antlr.Token { return s.op }

func (s *BitShiftContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitShiftContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterBitShift(s)
	}
}

func (s *BitShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitBitShift(s)
	}
}

func (s *BitShiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitBitShift(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionGroupContext struct {
	*ExpressionContext
}

func NewExpressionGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionGroupContext {
	var p = new(ExpressionGroupContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionGroupContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterExpressionGroup(s)
	}
}

func (s *ExpressionGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitExpressionGroup(s)
	}
}

func (s *ExpressionGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitExpressionGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewRelationalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalContext {
	var p = new(RelationalContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RelationalContext) GetOp() antlr.Token { return s.op }

func (s *RelationalContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RelationalContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterRelational(s)
	}
}

func (s *RelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitRelational(s)
	}
}

func (s *RelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ExpressionContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) CHARS() antlr.TerminalNode {
	return s.GetToken(FloParserCHARS, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerContext struct {
	*ExpressionContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(FloParserINTEGER, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetItemContext struct {
	*ExpressionContext
}

func NewGetItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetItemContext {
	var p = new(GetItemContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GetItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetItemContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GetItemContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetItemContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(FloParserL_BRACKET, 0)
}

func (s *GetItemContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(FloParserR_BRACKET, 0)
}

func (s *GetItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterGetItem(s)
	}
}

func (s *GetItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitGetItem(s)
	}
}

func (s *GetItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitGetItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	*ExpressionContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOATING_POINT() antlr.TerminalNode {
	return s.GetToken(FloParserFLOATING_POINT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContext struct {
	*ExpressionContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(FloParserNOT, 0)
}

func (s *NotContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitwiseOrContext struct {
	*ExpressionContext
}

func NewBitwiseOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseOrContext {
	var p = new(BitwiseOrContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BitwiseOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseOrContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitwiseOrContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitwiseOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterBitwiseOr(s)
	}
}

func (s *BitwiseOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitBitwiseOr(s)
	}
}

func (s *BitwiseOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitBitwiseOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	*ExpressionContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(FloParserLOGICAL_AND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitwiseAndContext struct {
	*ExpressionContext
}

func NewBitwiseAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseAndContext {
	var p = new(BitwiseAndContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BitwiseAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseAndContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitwiseAndContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitwiseAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterBitwiseAnd(s)
	}
}

func (s *BitwiseAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitBitwiseAnd(s)
	}
}

func (s *BitwiseAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitBitwiseAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	*ExpressionContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(FloParserL_BRACKET, 0)
}

func (s *ListContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(FloParserR_BRACKET, 0)
}

func (s *ListContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type XORContext struct {
	*ExpressionContext
}

func NewXORContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XORContext {
	var p = new(XORContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *XORContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XORContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *XORContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *XORContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterXOR(s)
	}
}

func (s *XORContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitXOR(s)
	}
}

func (s *XORContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitXOR(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	*ExpressionContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(FloParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(FloParserFALSE, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpressionContext struct {
	*ExpressionContext
}

func NewCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpressionContext {
	var p = new(CallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CallExpressionContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *CallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterCallExpression(s)
	}
}

func (s *CallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitCallExpression(s)
	}
}

func (s *CallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerContext struct {
	*ExpressionContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowerContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FloParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, FloParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(134)
			p.Match(FloParserL_PAREN)
		}
		{
			p.SetState(135)
			p.expression(0)
		}
		{
			p.SetState(136)
			p.Match(FloParserR_PAREN)
		}

	case 2:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(138)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UnaryContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == FloParserADDITION || _la == FloParserSUBTRACTION) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UnaryContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(139)
			p.expression(20)
		}

	case 3:
		localctx = NewUnaryIncDecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(140)
			p.Match(FloParserIDENTIFIER)
		}
		p.SetState(141)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UnaryIncDecContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == FloParserPLUS_PLUS || _la == FloParserMINUS_MINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UnaryIncDecContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 4:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(142)
			p.Match(FloParserNOT)
		}
		{
			p.SetState(143)
			p.expression(10)
		}

	case 5:
		localctx = NewListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(FloParserL_BRACKET)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
			{
				p.SetState(145)
				p.Expression_list()
			}

		}
		{
			p.SetState(148)
			p.Match(FloParserR_BRACKET)
		}

	case 6:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.Match(FloParserINTEGER)
		}

	case 7:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.Match(FloParserFLOATING_POINT)
		}

	case 8:
		localctx = NewReadIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(151)
			p.Match(FloParserIDENTIFIER)
		}

	case 9:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(FloParserTRUE)
		}

	case 10:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(FloParserFALSE)
		}

	case 11:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Match(FloParserCHARS)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(158)
					p.Match(FloParserPOWER)
				}
				{
					p.SetState(159)
					p.expression(22)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				p.SetState(161)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MulDivContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserMOD)|(1<<FloParserMULTIPLY)|(1<<FloParserDIVIDE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MulDivContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(162)
					p.expression(18)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				p.SetState(164)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == FloParserADDITION || _la == FloParserSUBTRACTION) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AddSubContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(165)
					p.expression(17)
				}

			case 4:
				localctx = NewBitShiftContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				p.SetState(167)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BitShiftContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == FloParserLSHIFT || _la == FloParserRSHIFT) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BitShiftContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(168)
					p.expression(16)
				}

			case 5:
				localctx = NewBitwiseAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(170)
					p.Match(FloParserT__2)
				}
				{
					p.SetState(171)
					p.expression(15)
				}

			case 6:
				localctx = NewXORContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(173)
					p.Match(FloParserT__3)
				}
				{
					p.SetState(174)
					p.expression(14)
				}

			case 7:
				localctx = NewBitwiseOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(176)
					p.Match(FloParserOR)
				}
				{
					p.SetState(177)
					p.expression(13)
				}

			case 8:
				localctx = NewRelationalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(179)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RelationalContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserEQUALS)|(1<<FloParserNOT_EQUALS)|(1<<FloParserLESS)|(1<<FloParserLESS_OR_EQUALS)|(1<<FloParserGREATER)|(1<<FloParserGREATER_OR_EQUALS))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RelationalContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(180)
					p.expression(12)
				}

			case 9:
				localctx = NewAndContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(182)
					p.Match(FloParserLOGICAL_AND)
				}
				{
					p.SetState(183)
					p.expression(10)
				}

			case 10:
				localctx = NewOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(185)
					p.Match(FloParserLOGICAL_OR)
				}
				{
					p.SetState(186)
					p.expression(9)
				}

			case 11:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(188)
					p.Match(FloParserL_PAREN)
				}
				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FloParserTRUE)|(1<<FloParserFALSE)|(1<<FloParserNOT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(FloParserADDITION-33))|(1<<(FloParserSUBTRACTION-33))|(1<<(FloParserL_BRACKET-33))|(1<<(FloParserL_PAREN-33))|(1<<(FloParserINTEGER-33))|(1<<(FloParserCHARS-33))|(1<<(FloParserFLOATING_POINT-33))|(1<<(FloParserIDENTIFIER-33)))) != 0) {
					{
						p.SetState(189)
						p.Expression_list()
					}

				}
				{
					p.SetState(192)
					p.Match(FloParserR_PAREN)
				}

			case 12:
				localctx = NewGetItemContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FloParserRULE_expression)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(194)
					p.Match(FloParserL_BRACKET)
				}
				{
					p.SetState(195)
					p.expression(0)
				}
				{
					p.SetState(196)
					p.Match(FloParserR_BRACKET)
				}

			}

		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_expression_list
	return p
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (s *Expression_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitExpression_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Expression_list() (localctx IExpression_listContext) {
	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FloParserRULE_expression_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.expression(0)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FloParserT__1 {
		{
			p.SetState(204)
			p.Match(FloParserT__1)
		}
		{
			p.SetState(205)
			p.expression(0)
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FloParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FloParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(FloParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FloVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FloParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FloParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(FloParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(FloParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(213)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(214)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "p.checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

func (p *FloParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 16:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FloParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FloParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.lineTerminatorAhead()

	case 13:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
