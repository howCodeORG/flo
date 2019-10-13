package lexer

import (
	"flo/token"
	"strings"
)

type Lexer struct {
	input      string
	SplitInput []string

	currentPosition int
	nextPosition    int
	ch              byte
	lastToken       token.Token
	CurrentLine     int
	CurrentColumn   int
}

func New(input string) *Lexer {

	l := &Lexer{
		input:         input,
		SplitInput:    strings.Split(input, "\n"),
		CurrentLine:   1,
		CurrentColumn: 0,
	}
	l.readCharacter()
	return l

}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token
checkComments:
	l.skipWhitespace()

	// Remove comments first
	switch l.ch {
	case '/':
		if l.peekCharacter() == '/' {
			// Found a single-line comment
			l.readSLComment()
		} else if l.peekCharacter() == '*' {
			// Found a multi-line comment
			l.readMLComment()
		} else {
			break
		}
		// If eat up an entire comment, we need to check again for comments
		// to ensure we don't have another comment directly after the first.
		// Without this goto, the code will move on and will miss the second
		// comment if it's directly after the first comment with no new tokens
		// in-between.
		goto checkComments
	}

	switch l.ch {
	default:
		if isDigit(l.ch) {
			tok.Value, tok.Type = l.readNumber()
			tok.Line = l.CurrentLine
			tok.Column = l.CurrentColumn
			l.lastToken = tok
			return tok
		} else if isLetter(l.ch) {
			tok.Value, tok.Type = l.readIdentifier()
			tok.Line = l.CurrentLine
			tok.Column = l.CurrentColumn
			l.lastToken = tok
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	case '"':
		fallthrough
	case '\'':
		tok.Value, tok.Type = l.readString()
		tok.Line = l.CurrentLine
		tok.Column = l.CurrentColumn
	case '+':
		tok = l.newToken(token.PLUS, string(l.ch))
	case '-':
		tok = l.newToken(token.MINUS, string(l.ch))
	case '%':
		tok = l.newToken(token.MODULO, string(l.ch))
	case '*':
		if l.peekCharacter() == '*' {
			tok = l.newToken(token.POWER, string(l.ch)+string(l.ch))
			l.readCharacter()
		} else {
			tok = l.newToken(token.MULTIPLY, string(l.ch))
		}
	case '<':
		if l.peekCharacter() == '=' {
			tok = l.newToken(token.LTEQ, "<=")
			l.readCharacter()
		} else {
			tok = l.newToken(token.LT, string(l.ch))
		}
	case '>':
		if l.peekCharacter() == '=' {
			tok = l.newToken(token.GTEQ, ">=")
			l.readCharacter()
		} else {
			tok = l.newToken(token.GT, string(l.ch))
		}
	case '!':
		if l.peekCharacter() == '=' {
			tok = l.newToken(token.NTEQ, "!=")
			l.readCharacter()
		} else {
			tok = l.newToken(token.BANG, string(l.ch))
		}
	case '{':
		tok = l.newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = l.newToken(token.RBRACE, string(l.ch))
	case ',':
		tok = l.newToken(token.COMMA, string(l.ch))
	case '/':
		tok = l.newToken(token.DIVIDE, string(l.ch))
	case '(':
		tok = l.newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = l.newToken(token.RPAREN, string(l.ch))
	case '[':
		tok = l.newToken(token.LBRACK, string(l.ch))
	case ']':
		tok = l.newToken(token.RBRACK, string(l.ch))
	case '=':
		if l.peekCharacter() == '=' {
			tok = l.newToken(token.EQEQ, string(l.ch)+string(l.ch))
			l.readCharacter()
		} else {
			tok = l.newToken(token.ASSIGN, string(l.ch))
		}
	case ';':
		tok = l.newToken(token.SEMICOLON, string(l.ch))
	case 0:
		tok.Type = token.EOF
		tok.Value = ""
		tok.Line = l.CurrentLine
		tok.Column = l.CurrentColumn
	}

	l.readCharacter()
	l.lastToken = tok
	return tok
}

func (l *Lexer) readSLComment() {
	for l.ch != '\n' && l.ch != '\r' && l.ch != 0 {
		l.readCharacter()
	}
	l.skipWhitespace()
}

func (l *Lexer) readMLComment() {
	for l.ch != 0 {
		l.readCharacter()
		if l.ch == '*' && l.peekCharacter() == '/' {
			l.readCharacter()
			l.readCharacter()
			l.skipWhitespace()
			break
		}
	}
}

func (l *Lexer) skipWhitespace() {

	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		// Semicolon injection
		if l.ch == '\n' {
			switch l.lastToken.Type {
			case token.INT:
				fallthrough
			case token.FLOAT:
				fallthrough
			case token.IDENT:
				fallthrough
			case token.STRING:
				fallthrough
			case token.RPAREN:
				fallthrough
			case token.RBRACE:
				fallthrough
			case token.K_TRUE:
				fallthrough
			case token.K_FALSE:
				fallthrough
			case token.K_INT:
				fallthrough
			case token.K_FLOAT:
				fallthrough
			case token.K_STRING:
				fallthrough
			case token.K_BOOL:
				l.ch = ';'
				l.CurrentColumn = 0
				l.CurrentLine++
				return
			}
		}

		l.readCharacter()
	}

}

func (l *Lexer) readString() (string, string) {
	isDoubleQuoted := false
	if l.ch == '"' {
		isDoubleQuoted = true
		l.readCharacter()
	} else if l.ch == '\'' {
		isDoubleQuoted = false
		l.readCharacter()
	}
	// Read string until next quote
	if isDoubleQuoted {
		position := l.currentPosition
		for l.ch != '"' {
			if l.ch == '\n' || l.ch == '\r' || l.ch == 0 {
				return l.input[position:l.currentPosition], token.ILLEGAL
			}
			l.readCharacter()
		}
		return l.input[position:l.currentPosition], token.STRING
	} else {
		position := l.currentPosition
		for l.ch != '\'' {
			if l.ch == '\n' || l.ch == '\r' || l.ch == 0 {
				return l.input[position:l.currentPosition], token.ILLEGAL
			}
			l.readCharacter()
		}
		return l.input[position:l.currentPosition], token.STRING
	}
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (l *Lexer) readNumber() (string, string) {
	position := l.currentPosition

	for isDigit(l.ch) {
		l.readCharacter()
	}
	// Must be a float
	if l.ch == '.' {
		l.readCharacter()
	} else {
		// Must be an int
		return l.input[position:l.currentPosition], token.INT
	}

	// Get right-hand side of dot
	for isDigit(l.ch) {
		l.readCharacter()
	}

	return l.input[position:l.currentPosition], token.FLOAT
}

func isLetter(character byte) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (character == '_')
}

func isValidInIdentifier(character byte) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (character >= '0' && character <= '9') || (character == '_')
}

func (l *Lexer) readIdentifier() (string, string) {
	position := l.currentPosition

	// identifiers must start with a letter or an underscore
	for isLetter(l.ch) {
		l.readCharacter()
	}
	// after that they can contain numbers
	for isValidInIdentifier(l.ch) {
		l.readCharacter()
	}

	// Check if identifier is a keyword
	string := l.input[position:l.currentPosition]
	tokenType := token.IDENT

	tokenType = l.getKeyword(string)

	return string, tokenType
}

func (l *Lexer) readDigit() string {
	position := l.currentPosition

	for isDigit(l.ch) {
		l.readCharacter()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) newToken(tokenType string, tokenValue string) token.Token {
	return token.Token{Type: tokenType, Value: tokenValue, Line: l.CurrentLine, Column: l.CurrentColumn}
}

func (l *Lexer) peekCharacter() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readCharacter() {

	if l.ch == '\n' {
		l.CurrentLine++
		l.CurrentColumn = 0
	}
	if l.ch != 0 {
		l.CurrentColumn++
	}

	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition++

}

func (l *Lexer) getKeyword(keyword string) string {
	if val, ok := token.Keywords["k_"+keyword]; ok {
		return val
	}
	return token.IDENT
}
