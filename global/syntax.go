package global

import "github.com/emirpasic/gods/sets/linkedhashset"

// Syntax should define how to go from Token -> AST and AST -> Token
type Syntax struct {
	tokenTypes            *linkedhashset.Set
	commentPrefix         string
	multilineCommentStart string
	multilineCommentEnd   string
	stringWrapper         rune
	charWrapper           rune
}
