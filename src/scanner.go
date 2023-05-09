package main

import (
	"bufio"
	"errors"
	"fmt"
	"unicode"
)

type Scanner struct {
	lineCount, colCount int
	needToken           bool
	lastToken           SuperToken
	inStream            *bufio.Reader
}

func (scanner *Scanner) putBackToken() {
	scanner.needToken = false
}

func (scanner *Scanner) getToken() (*SuperToken, error) {
	if !scanner.needToken {
		scanner.needToken = true
		return &scanner.lastToken, nil
	}
	var t SuperToken
	state := 0
	foundOne := false
	var c byte
	var lex string
	var TokenType string
	var column, line int

	c, err := scanner.inStream.ReadByte()
	if err != nil {
		foundOne = true
		TokenType = "EOF"
	}
	for !foundOne {
		scanner.colCount++
		switch state {
		case 0:
			lex = ""
			column = scanner.colCount
			line = scanner.lineCount
			cs := string(c)
			if unicode.IsLetter(rune(c)) {
				state = 1
			} else if unicode.IsDigit(rune(c)) {
				state = 2
			} else if cs == "+" {
				state = 3
			} else if cs == "-" {
				state = 4
			} else if cs == "*" {
				state = 5
			} else if cs == "/" {
				state = 6
			} else if cs == "(" {
				state = 7
			} else if cs == ")" {
				state = 8
			} else if cs == "<" {
				state = 9
			} else if cs == ">" {
				state = 11
			} else if cs == "=" {
				state = 13
			} else if cs == "%" {
				state = 14
			} else if cs == "^" {
				state = 15
			} else if cs == "?" {
				state = 16
			} else if c == '\n' {
				scanner.colCount = -1
				scanner.lineCount++
			} else if c == '\r' {

			} else if c == 0 {
				foundOne = true
				TokenType = "EOF"
			} else if unicode.IsSpace(rune(c)) {

			} else {
				fmt.Print("Unrecognized Token found at line ")
				fmt.Print(line)
				return nil, errors.New("Unrecognized Token")
			}
			break
		case 1:
			if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
				state = 1
			} else {
				if lex == "set" {
					foundOne = true
					TokenType = "keyword"
				}
				if !foundOne {
					TokenType = "identifier"
					foundOne = true
				}
			}
			break
		case 2:
			if unicode.IsDigit(rune(c)) {
				state = 2
			} else {
				TokenType = "number"
				foundOne = true
			}
			break
		case 3:
			TokenType = "add"
			foundOne = true
			break
		case 4:
			TokenType = "sub"
			foundOne = true
			break
		case 5:
			TokenType = "times"
			foundOne = true
			break
		case 6:
			TokenType = "divide"
			foundOne = true
			break
		case 7:
			TokenType = "lparen"
			foundOne = true
			break
		case 8:
			TokenType = "rparen"
			foundOne = true
			break
		case 9:
			if c == '<' {
				state = 10
			} else {
				fmt.Println("se esperaba un < en la linea " + string(line) + "y columna" + string(column))
			}
			break
		case 10:
			TokenType = "lbyte"
			foundOne = true
			break
		case 11:
			if c == '>' {
				state = 12
			} else {
				fmt.Println("se esperaba un < en la linea " + string(line) + "y columna" + string(column))
			}
			break
		case 12:
			TokenType = "rbyte"
			foundOne = true
			break
		case 13:
			TokenType = "equals"
			foundOne = true
			break
		case 14:
			TokenType = "mod"
			foundOne = true
			break
		case 15:
			TokenType = "pot"
			foundOne = true
			break
		case 16:
			TokenType = "sqrt"
			foundOne = true
			break
		}
		if !foundOne {
			lex = lex + string(c)
			c, err = scanner.inStream.ReadByte()
		}
	}
	scanner.inStream.UnreadByte()
	scanner.colCount--
	if TokenType == "number" || TokenType == "identifier" || TokenType == "keyword" {
		lexString := string(lex)
		lt := LexicalToken{
			typ:    TokenType,
			lexeme: lexString,
			line:   line,
			col:    column,
		}
		t = &lt
	} else {
		lt := Token{
			typ:  TokenType,
			line: line,
			col:  column,
		}
		t = &lt
	}
	scanner.lastToken = t
	return &t, nil
}

/*
; */
