package common

import (
	"github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/markphelps/optional"
)

// Syntax should define how to go from Token -> AST and AST -> Token
type Syntax struct {
	tokenTypes    *linkedhashset.Set
	commentPrefix optional.String

	stringWrapper rune
	charWrapper   rune
}

type MultiLineCommentSyntax struct {
	multilineCommentStart string
	multilineCommentEnd   string
}
