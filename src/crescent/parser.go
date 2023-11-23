package crescent

import (
	"crescentLang/common"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/moznion/go-optional"
	"go/ast"
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
			},
			Parser: func(p *common.Parser, i []common.Token) (common.Node, error) {
				p.
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
