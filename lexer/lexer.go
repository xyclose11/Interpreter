package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current pos in input (aka current char)
	readPosition int  // next char in input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // initialize
	l.readChar()
	return l
}

// get next character and advance through input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL
	} else {
		l.ch = l.input[l.readPosition] // access next character
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '*':
		tok = newToken(token.ASTERISK, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch                                              // Get current char
			l.readChar()                                            // Increment to get next char
			literal := string(ch) + string(l.ch)                    // Create new literal
			tok = token.Token{Type: token.NOT_EQ, Literal: literal} // Create token
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	case '/':
		tok = newToken(token.SLASH, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Similar to readChar except it does not increment l at all
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// IP: readChar() add UTF-8/UNICODE support
// readNumber() support floats, hex notation, octal, etc
// create method 'makeTwoCharToken' that will peek and advance if it found the correct token PAGE 26
