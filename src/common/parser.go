package common

import (
	"github.com/emirpasic/gods/sets"
	"go/ast"
)

// Pattern is an array of a set of TokenType
type Pattern = []sets.Set

type Parser struct {
	nodes  []Node
	index  uint
	syntax Syntax
}

type Scope struct {
	pattern    Pattern
	converters []Converter
}

type Converter struct {
	block   func(parser *Parser)
	pattern Pattern
}

func (p *Parser) Parse(tokens []Token) []ast.Node {

	/*
		for _, token := range tokens {
			token.
		}

		ast.Node()
	*/
	return nil
}
