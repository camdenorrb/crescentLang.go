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
	Lex(lines chan string) []Token
}

type GenericLexer struct {
	reader *bufio.Reader
	syntax *Syntax
	pos    Position
}

func NewGenericLexer(reader io.Reader, syntax *Syntax) *GenericLexer {
	return &GenericLexer{
		reader: bufio.NewReader(reader),
		syntax: syntax,
		pos:    Position{line: 1, column: 0},
	}
}

// Plan: If syntax doesn't contain Rune, append and check again
// Make sure to check if number, string literal, etc

func (l *GenericLexer) lexLine(line string, lineNumber uint) ([]Token, error) {

	var tokens []Token
	var number string
	var builder string

	for index, character := range []rune(line) {

		wasNumber := false

		if builder != "" {
			if tokenType, exists := l.syntax.tokenTypes[builder]; exists {

				tokens = append(tokens, Token{
					ColumnRange: IntRange{Start: index - len(builder), End: index},
					LineNumber:  lineNumber,
					Type:        tokenType,
				})

				builder = ""
			} else if !l.syntax.identifierCharacterValidator(character) {

				// TODO: Error with line and column info for unknown keyword
				tokens = append(tokens, Token{
					Value:       builder,
					ColumnRange: IntRange{Start: index - len(builder), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.identifierTokenType,
				})

				builder = ""
			} else {
				builder += string(character)
				continue
			}
		}

		switch character {

		case '"':

		case '\'':

		case '.':
			if len(number) > 0 {
				number += "."
				wasNumber = true
			}

		// Note: Negative numbers won't be lexed entirely, will be Minus token first
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

			// If there is a hanging dot, remove token and add to number
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == l.syntax.tokenTypes["."] {
				tokens = tokens[:len(tokens)-1]
				number += "."
			}

			number += string(character)
			wasNumber = true

		default:
			builder += string(character)

		}

		if !wasNumber && number != "" {
			// TODO: Store number as token
			number = ""
		}

	}

	// Must be identifier
	if builder != "" {

		tokens = append(tokens, Token{
			Value:       builder,
			ColumnRange: IntRange{Start: len(line) - len(builder) - 1, End: len(line) - 1},
			LineNumber:  lineNumber,
			Type:        l.syntax.identifierTokenType,
		})

		builder = ""
	}

	return tokens, nil
}

/*

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
*/
