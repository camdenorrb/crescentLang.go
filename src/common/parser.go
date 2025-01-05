package common

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/joomcode/errorx"
)

//type ParserType uint

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
	Syntax   *Syntax
	Patterns []Pattern
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

	matchingPatterns := p.Patterns

	var nodes []Node
	var tokenCache []Token
	var index int

	for _, token := range tokens {

		tokenCache = append(tokenCache, token)

		var validPatterns []Pattern

		for _, pattern := range matchingPatterns {

			// If the pattern is longer than the token cache or the token type is not valid for the pattern
			if len(pattern.Tokens) <= index || !pattern.Tokens[index].Contains(token.Type) {
				continue
			}

			validPatterns = append(validPatterns, pattern)
		}

		matchingPatterns = validPatterns

		// If only one pattern matches and the pattern is complete
		if len(matchingPatterns) == 1 && len(matchingPatterns[0].Tokens) == len(tokenCache) {

			node, err := matchingPatterns[0].Parser(p, tokenCache)
			if err != nil {
				return nil, errorx.IllegalState.Wrap(err, "Failed to parse")
			}

			nodes = append(nodes, node)

			matchingPatterns = p.Patterns // Reset
			tokenCache = tokenCache[:0]   // Clear
		}

		if len(matchingPatterns) == 0 {
			return nil, errorx.IllegalState.New("No matching patterns for %+v", tokenCache)
		}

		index++
	}

	// If there are multiple matching patterns or there is only one matching pattern but the pattern is not complete
	if (len(matchingPatterns) > 1 && len(tokenCache) > 1) || (len(matchingPatterns) == 1 && len(matchingPatterns[0].Tokens) != len(tokenCache)) {
		return nil, errorx.IllegalState.New("Ambiguous patterns for %+v", tokenCache)
	}

	// If there is only one matching pattern and the pattern is complete
	if len(matchingPatterns) == 1 && len(matchingPatterns[0].Tokens) == len(tokenCache) {

		node, err := matchingPatterns[0].Parser(p, tokenCache)
		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "Failed to parse")
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
