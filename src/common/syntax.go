package common

import (
	"github.com/moznion/go-optional"
)

type KeyTokenTypes struct {
	IdentifierTokenType TokenType
	CharTokenType       TokenType
	StringTokenType     TokenType
}

// Syntax should define how to go from Token -> AST and AST -> Token
type Syntax struct {
	KeyTokenTypes
	tokenTypes             map[string]TokenType
	commentPrefix          optional.Option[string]
	multiLineCommentSyntax optional.Option[MultiLineCommentSyntax]
}

type MultiLineCommentSyntax struct {
	multilineCommentStart string
	multilineCommentEnd   string
}
