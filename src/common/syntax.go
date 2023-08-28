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
	TokenTypes             map[string]TokenType
	CommentPrefix          optional.Option[string]
	MultiLineCommentSyntax optional.Option[MultiLineCommentSyntax]
}

type MultiLineCommentSyntax struct {
	MultilineCommentStart string
	MultilineCommentEnd   string
}
