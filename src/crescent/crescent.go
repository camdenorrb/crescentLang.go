package crescent

import (
	"crescentLang/common"
	"go/ast"
)

func Lex() {

}

func Parse(tokens []common.Token) ([]ast.Node, error) {
	return parser.Parse(tokens)
}
