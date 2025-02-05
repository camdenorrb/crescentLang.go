package crescent

import (
	"crescentLang/common"
	mapset "github.com/deckarep/golang-set/v2"
)

// TODO: Maybe a parser combinator (Bunch of small parsers making a bigger one)

const (
	UnsetParserMode common.ParserMode = iota
	FunctionParameterParserMode
	FunctionBodyParserMode
)

var Patterns = map[common.ParserMode][]common.Pattern{
	UnsetParserMode: {
		{
			Name: "Function",
			Tokens: []mapset.Set[common.TokenType]{
				mapset.NewSet(FUN),
				mapset.NewSet(IDENTIFIER),
				mapset.NewSet(CURLYL),
				mapset.NewSet(CURLYR),
				// State switches to FunctionBodyParserMode
			},
			Parser: func(p *common.Parser, tokens []common.Token) (common.Node, common.ParserMode, error) {
				return common.Function{
					Name:       tokens[1].Value.(string),
					Visibility: common.Visibility(tokens[0].Value.(string)),
					Parameters: nil,
					Body:       nil,
					ReturnType: nil,
				}, FunctionBodyParserMode, nil
			},
		},
		{
			Name: "Function With Visibility",
			Tokens: []mapset.Set[common.TokenType]{
				mapset.NewSet(PUBLIC, PRIVATE, INTERNAL),
				mapset.NewSet(FUN),
				mapset.NewSet(IDENTIFIER),
				mapset.NewSet(PARENL),
				mapset.NewSet(PARENR),
				mapset.NewSet(CURLYL),
				// State switches to FunctionBodyParserMode
			},
			Parser: func(p *common.Parser, tokens []common.Token) (common.Node, common.ParserMode, error) {
				return common.Function{
					Name:       tokens[1].Value.(string),
					Visibility: common.Visibility(tokens[0].Value.(string)),
					Parameters: nil,
					Body:       nil,
					ReturnType: nil,
				}, FunctionBodyParserMode, nil
			},
		},
	},
	FunctionBodyParserMode: {
		{
			Name: "VariableWithAssignment",
			Tokens: []mapset.Set[common.TokenType]{
				mapset.NewSet(CONST, VAL, VAR),
				mapset.NewSet(IDENTIFIER),
				mapset.NewSet(ASSIGN),
				mapset.NewSet(STRING, NUMBER, CHAR),
			},
		},
		{
			Name: "For Loop",
			Tokens: []mapset.Set[common.TokenType]{
				mapset.NewSet(FOR),
				mapset.NewSet(IDENTIFIER),
			},
		},
	},
}

/*
var patterns = map[common.ParserMode][]common.Pattern{
	UnsetParserMode: {
		(expressionPatterns...),
	},
}
*/

var parser = common.Parser{
	Syntax:   Syntax,
	Patterns: Patterns,
}
