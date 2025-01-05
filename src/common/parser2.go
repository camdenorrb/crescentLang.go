package common

import mapset "github.com/deckarep/golang-set/v2"

type Patterns []mapset.Set[TokenType]

func (p Patterns) StartsWith(tokens []Token) bool {

	if len(tokens) > len(p) {
		return false
	}

	for index, pattern := range p[:len(tokens)] {
		if !pattern.Contains(tokens[index].Type) {
			return false
		}
	}

	return true
}

func (p Patterns) Matches(tokens []Token) bool {

	if len(p) != len(tokens) {
		return false
	}

	for index, pattern := range p {
		if !pattern.Contains(tokens[index].Type) {
			return false
		}
	}

	return true
}

type FeatureParser struct {
	Feature             string
	AllowedInBody       []FeatureParser
	IsSelfAllowedInBody bool
	EntryPatterns       Patterns
	ExitPatterns        Patterns
	// Parser takes in only the matching tokens for entry and exit patterns
	// Nodes are from what was allowed in the body
	// 	Parser func(Entry: []Token, Exit: []Token, Body: []Node) (Node, error)
	Parser func([]Token, []Token, []Node) (Node, error)
}

type Parser2 struct {
	Syntax         *Syntax
	FeatureParsers []FeatureParser
}

func (p *Parser2) Parse(tokens []Token) ([]Node, error) {

	var nodes []Node
	var tokenCache []Token

	for _, token := range tokens {

		tokenCache = append(tokenCache, token)

		for _, featureParser := range p.FeatureParsers {
			featureParser.EntryPatterns.StartsWith(tokenCache)
		}

		if len(validPatterns) == 1 && len(validPatterns[0].Tokens) == len(tokenCache) {

			node, err := validPatterns[0].Parser(tokenCache, nil, nil)
			if err != nil {
				return nil, err
			}

			nodes = append(nodes, node)
			tokenCache = nil
		}

		if len(validPatterns) == 0 {
			return nil, IllegalState.New("No matching patterns for %+v", tokenCache)
		}
	}

	return nil, nil
}
