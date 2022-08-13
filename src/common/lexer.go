package common

import (
	"bufio"
	"io"
)

type Position struct {
	line   uint
	column uint
}

type Lexer interface {
	Lex() []Token
}

type GenericLexer struct {
	reader *bufio.Reader
	syntax *Syntax
	pos    Position
}

func NewGenericLexer(reader io.Reader) *GenericLexer {
	return &GenericLexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Plan: If syntax doesn't contain Rune, append and check again
// Make sure to check if number, string literal, etc

func (l *GenericLexer) Lex() ([]Token, error) {

	builder := ""

	for {

		rune, _, err := l.reader.ReadRune()
		if err != nil {

			if err != io.EOF {
				return nil, err
			}

			break
		}

	}
}
