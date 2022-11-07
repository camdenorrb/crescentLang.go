package crescent

import (
	"crescentLang/common"
	mapset "github.com/deckarep/golang-set/v2"
)

const (
	UnsetParserMode common.ParserMode = iota
	ExpressionParserMode
)

var expressionPatterns = []common.Pattern{
	{
		Name: "Variable",
		Tokens: []mapset.Set[common.TokenType]{
			mapset.NewSet(CONST, VAL, VAR),
			mapset.NewSet(IDENTIFIER),
			mapset.NewSet(EQUALS),
		},
	},
	{
		Name: "For Loop",
		Tokens: []mapset.Set[common.TokenType]{
			mapset.NewSet(FOR),
			mapset.NewSet(IDENTIFIER),
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

var parser = common.Parser{}
