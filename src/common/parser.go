package common

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/joomcode/errorx"
)

type ParserMode uint

/*
const (
	UnsetParserType ParserType = iota
	ScopeParserType
	ExprParserType
	StatementParserType
)
*/

// Pattern is an array of a set of TokenType
// Meaning that for each index there can be multiple valid token types
type Pattern struct {
	Name   string
	Parser func(*Parser, []Token) (Node, error) // Can be nil, if nil then the node is automatically generated
	Tokens []mapset.Set[TokenType]
}

type Parser struct {
	Syntax           *Syntax
	Patterns         map[ParserMode][]Pattern
	Mode             ParserMode
	matchingPatterns []Pattern // TODO: Inline
	tokenCache       []Token   // TODO: Inline
	index            int       // TODO: Inline
}

/*
type Scope struct {
	pattern Pattern
}
*/

/*
type Converter struct {
	block   func(Parser *Parser)
	pattern Pattern
}
*/

// The registered pattern parsers should control the Parser state

// TODO: Error result should have the line and column number
func (p *Parser) Parse(tokens []Token) ([]Node, error) {

	var nodes []Node

	p.matchingPatterns = p.Patterns[p.Mode]

	for _, token := range tokens {

		p.tokenCache = append(p.tokenCache, token)

		var validPatterns []Pattern

		for _, pattern := range p.matchingPatterns {

			if len(pattern.Tokens) <= p.index || !pattern.Tokens[p.index].Contains(token.Type) {
				continue
			}

			validPatterns = append(validPatterns, pattern)
		}

		p.matchingPatterns = validPatterns

		if len(p.matchingPatterns) == 1 && len(p.matchingPatterns[0].Tokens) == len(p.tokenCache) {

			node, err := p.matchingPatterns[0].Parser(p, p.tokenCache)
			if err != nil {
				return nil, err
			}

			nodes = append(nodes, node)
			p.matchingPatterns = p.Patterns[p.Mode]
			p.tokenCache = nil
		}

		if len(p.matchingPatterns) == 0 {
			return nil, errorx.IllegalState.New("No matching patterns for %+v", p.tokenCache)
		}

		p.index++
	}

	if (len(p.matchingPatterns) > 1 && len(p.tokenCache) > 1) || (len(p.matchingPatterns) == 1 && len(p.matchingPatterns[0].Tokens) != len(p.tokenCache)) {
		return nil, errorx.IllegalState.New("Ambiguous patterns for %+v", p.tokenCache)
	}

	if len(p.matchingPatterns) == 1 {

		pattern := p.matchingPatterns[0]

		node, err := pattern.Parser(p, p.tokenCache)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
