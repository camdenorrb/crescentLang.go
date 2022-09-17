package common

import (
	"github.com/moznion/go-optional"
)

type KeyTokenTypes struct {
	IdentifierTokenType       optional.Option[TokenType]
	CharTokenType             optional.Option[TokenType]
	StringTokenType           optional.Option[TokenType]
	NumberTokenType           optional.Option[TokenType]
	CommentTokenType          optional.Option[TokenType]
	MultiLineCommentTokenType optional.Option[TokenType]
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
