package token

type Token struct {
	Type   string
	Value  string
	Line   int
	Column int
}

var Keywords = map[string]string{
	K_IF:     K_IF,
	K_ELSE:   K_ELSE,
	K_INT:    K_INT,
	K_STRING: K_STRING,
	K_FLOAT:  K_FLOAT,
	K_BOOL:   K_BOOL,
	K_TRUE:   K_TRUE,
	K_NOT:    K_NOT,
	K_FALSE:  K_FALSE,
	K_PRINT:  K_PRINT,
	K_FOR:    K_FOR,
	K_FUNC:   K_FUNC,
	K_RETURN: K_RETURN,
	K_DYN:    K_DYN,
}

const (
	K_IF     = "k_if"
	K_ELSE   = "k_else"
	K_INT    = "k_int"
	K_STRING = "k_string"
	K_FLOAT  = "k_float"
	K_BOOL   = "k_bool"
	K_TRUE   = "k_true"
	K_NOT    = "k_not"
	K_FALSE  = "k_false"
	K_PRINT  = "k_print"
	K_FOR    = "k_for"
	K_FUNC   = "k_func"
	K_RETURN = "k_return"
	K_DYN    = "k_dyn"

	EOF       = ""
	NIL       = ""
	ILLEGAL   = "ILLEGAL"
	ASSIGN    = "="
	LT        = "<"
	GT        = ">"
	NTEQ      = "!="
	EQEQ      = "=="
	LTEQ      = "<="
	GTEQ      = ">="
	PLUS      = "+"
	MINUS     = "-"
	MULTIPLY  = "*"
	MODULO    = "%"
	DIVIDE    = "/"
	BANG      = "!"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	SEMICOLON = ";"
	DOT       = "."
	COMMA     = ","
	POWER     = "**"
	INT       = "INT"
	FLOAT     = "FLOAT"
	STRING    = "STRING"
	BOOL      = "BOOL"
	IDENT     = "IDENT"
)
