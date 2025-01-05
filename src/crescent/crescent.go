package crescent

import (
	"crescentLang/common"
)

func Lex() {

}

func Parse(tokens []common.Token) ([]common.Node, error) {
	return parser.Parse(tokens)
}
