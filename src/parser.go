package main

import (
	"fmt"
	"os"
	"strconv"
)

type Parser struct {
	scan *Scanner
}

type Parserr interface {
	Parse() AST
	prog() AST
	bitExpr() AST
	restBitExpr(e AST)
	expr() AST
	restExpr(e AST) AST
	term() AST
	restTerm(t AST) AST
	pot() AST
	potTerm(p AST) AST
	factor()
}

func (par *Parser) factor() AST {
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() == "lparen" {
		result := par.bitExpr()
		taux, _ := par.scan.getToken()
		t := *taux
		if t.getType() != "rparen" {
			fmt.Println("Expected )")
			fmt.Println("Program Aborted due to exception!")
			os.Exit(1)
		}
		return result
	} else if t.getType() == "sqrt" {
		e := par.bitExpr()
		return par.potTerm(&SqrtNode{e})
	} else {
		lex := t.getLex()
		val, _ := strconv.ParseFloat(lex, 64)
		return NumNode{val}
	}
}

func (par *Parser) potTerm(e AST) AST {
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() == "pot" {
		return par.potTerm(&PotNode{e, par.factor()})
	}
	/* fmt.Println(t.getType())
	fmt.Println()
	if t.getType() == "sqrt" {
		fmt.Println(("here2"))

	} */
	par.scan.putBackToken()
	return e
}

func (par *Parser) pot() AST {
	return par.potTerm(par.factor())
}

func (par *Parser) restTerm(e AST) AST {
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() == "times" {
		return par.restTerm(&TimesNode{e, par.pot()})
	}
	if t.getType() == "divide" {
		return par.restTerm(&DivideNode{e, par.pot()})
	}
	if t.getType() == "mod" {
		return par.restTerm(&ModNode{e, par.pot()})
	}
	par.scan.putBackToken()
	return e
}

func (par *Parser) term() AST {
	return par.restTerm(par.pot())
}

func (par *Parser) restExpr(e AST) AST {
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() == "add" {
		return par.restExpr(&AddNode{e, par.term()})
	}
	if t.getType() == "sub" {
		return par.restExpr(&SubNode{e, par.term()})
	}
	par.scan.putBackToken()
	return e
}

func (par *Parser) expr() AST {
	return par.restExpr(par.term())
}

func (par *Parser) restBitExpr(e AST) AST {
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() == "lbyte" {
		return par.restBitExpr(&LbyteNode{e, par.expr()})
	}
	if t.getType() == "rbyte" {
		return par.restBitExpr(&RbyteNode{e, par.expr()})
	}
	par.scan.putBackToken()
	return e
}

func (par *Parser) bitExpr() AST {
	return par.restBitExpr(par.expr())
}

func (par *Parser) prog() AST {
	result := par.bitExpr()
	taux, _ := par.scan.getToken()
	t := *taux
	if t.getType() != "EOF" {
		col := t.getCol()
		fmt.Println("Syntax Error: Expected EOF, found token at column", col)
	}
	return result
}

func (par *Parser) Parse() AST {
	return par.prog()
}
