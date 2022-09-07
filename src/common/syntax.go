package common

import (
	"github.com/markphelps/optional"
)

// Syntax should define how to go from Token -> AST and AST -> Token
type Syntax struct {
	identifierTokenType          TokenType
	identifierCharacterValidator func(rune) bool
	tokenTypes                   map[string]TokenType
	commentPrefix                optional.String
	stringWrapper                rune
	charWrapper                  rune
}

type MultiLineCommentSyntax struct {
	multilineCommentStart string
	multilineCommentEnd   string
}
