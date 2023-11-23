package main

import (
	"crescentLang/common"
	"crescentLang/language"
	"fmt"
	"github.com/moznion/go-optional"
	"os"
	"strings"
)

func main() {

	syntax := &common.Syntax{
		KeyTokenTypes: common.KeyTokenTypes{
			IdentifierTokenType:       optional.Some(common.TokenType(1)),
			CharTokenType:             optional.Some(common.TokenType(2)),
			StringTokenType:           optional.Some(common.TokenType(3)),
			NumberTokenType:           optional.Some(common.TokenType(4)),
			CommentTokenType:          optional.Some(common.TokenType(5)),
			MultiLineCommentTokenType: optional.Some(common.TokenType(6)),
		},
		TokenTypes: map[string]common.TokenType{
			"{": common.TokenType(7),
			"}": common.TokenType(8),
			"`": common.TokenType(9),
			"|": common.TokenType(10),
			":": common.TokenType(11),
		},
		CommentPrefix:          optional.Some("#"),
		MultiLineCommentSyntax: nil,
	}

	lexer, err := language.NewGenericLexer(syntax)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.ReadFile("language.lang")
	if err != nil {
		return
	}

	lex, err := lexer.Lex(strings.NewReader(string(file)))
	if err != nil {
		return
	}

	for _, token := range lex {
		println(token.Type)
	}

}
