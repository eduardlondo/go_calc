package main

import (
	"bufio"
	"strings"
)

func createScanner(str string) Scanner {
	r := strings.NewReader(str)
	reader := bufio.NewReader(r)
	scanner := Scanner{lineCount: 1, colCount: -1, needToken: true, lastToken: nil, inStream: reader}
	return scanner
}

func eval(expr string) int {
	scanner := createScanner(expr)
	parser := Parser{scan: &scanner}
	tree := parser.Parse()
	result := tree.evaluate()
	return result
}
