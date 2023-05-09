package main

type SuperToken interface {
	getType() string
	getLine() int
	getCol() int
	getLex() string
}

type Token struct {
	typ       string
	line, col int
}

type LexicalToken struct {
	typ, lexeme string
	line, col   int
}

func (token *Token) getType() string {
	return token.typ
}

func (token *Token) getLine() int {
	return token.line
}

func (token *Token) getCol() int {
	return token.col
}

func (token *Token) getLex() string {
	str := ""
	return str
}

func (token *LexicalToken) getType() string {
	return token.typ
}

func (token *LexicalToken) getLine() int {
	return token.line
}

func (token *LexicalToken) getCol() int {
	return token.col
}

func (token *LexicalToken) getLex() string {
	return token.lexeme
}
