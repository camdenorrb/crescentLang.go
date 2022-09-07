package common

import (
	"bufio"
	"io"
	"unicode"
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

type BuilderType uint

const (
	UnsetBuilderType BuilderType = iota
	KeywordBuilderType
	SymbolBuilderType
	StringBuilderType
	CharBuilderType
	NumberBuilderType
	CommentBuilderType
	MultiLineCommentBuilderType
)

// Plan: If syntax doesn't contain Rune, append and check again
// Make sure to check if number, string literal, etc

func (l *GenericLexer) lexLine(line string, lineNumber uint) ([]Token, error) {

	var tokens []Token

	var builder string
	var builderType BuilderType

	for index, character := range []rune(line) {

		wasNumber := false

		if builder != "" {
			// If matching token type is found, store token
			if tokenType, exists := l.syntax.tokenTypes[builder]; exists {

				tokens = append(tokens, Token{
					ColumnRange: IntRange{Start: index - len(builder), End: index},
					LineNumber:  lineNumber,
					Type:        tokenType,
				})

				builder = ""
				builderType = UnsetBuilderType
			} else if builderType == KeywordBuilderType {

				if unicode.IsLetter(character) {
					builder += string(character)
					continue
				}

				tokens = append(tokens, Token{
					Value:       builder,
					ColumnRange: IntRange{Start: index - len(builder), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.IdentifierTokenType,
				})

				builder = ""
			}
		}

		switch character {

		case '"':

			if builderType == StringBuilderType {

				tokens = append(tokens, Token{
					Value:       builder,
					ColumnRange: IntRange{Start: index - len(builder) - len("\""), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.StringTokenType,
				})

				builder = ""
				builderType = UnsetBuilderType
			} else {
				builderType = StringBuilderType
			}

		case '\'':
			if builderType == CharBuilderType {

				tokens = append(tokens, Token{
					Value:       builder,
					ColumnRange: IntRange{Start: index - len(builder) - len("'"), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.CharTokenType,
				})

				builder = ""
				builderType = UnsetBuilderType
			} else {
				builderType = CharBuilderType
			}

		case '.':
			if builderType != NumberBuilderType {
				tokens = append(tokens, Token{
					ColumnRange: IntRange{Start: index, End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.DotTokenType,
				})
			} else {
				builder += "."
				wasNumber = true
			}

		// Note: Negative numbers won't be lexed entirely, will be Minus token first
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

			// If there is a hanging dot, remove token and add to number
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == l.syntax.DotTokenType {
				tokens = tokens[:len(tokens)-1]
				builder += "."
			}

			builder += string(character)
			wasNumber = true

		default:
			if unicode.IsLetter(character) {
				builder += string(character)
				builderType = KeywordBuilderType
			} else {
				builder += string(character)
				builderType = SymbolBuilderType
			}
		}

		if builderType == NumberBuilderType && !wasNumber && builder != "" {
			// TODO: Store number as token
			builder = ""
			builderType = UnsetBuilderType
		}

	}

	// Must be identifier
	if builder != "" {

		switch builderType {
		// TODO: Switch between types
		}

		tokens = append(tokens, Token{
			Value:       builder,
			ColumnRange: IntRange{Start: len(line) - len(builder) - 1, End: len(line) - 1},
			LineNumber:  lineNumber,
			Type:        l.syntax.IdentifierTokenType,
		})

		builder = ""
		builderType = UnsetBuilderType
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
