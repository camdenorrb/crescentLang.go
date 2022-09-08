package common

import (
	"bufio"
	"io"
	"strings"
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

	var builder strings.Builder
	var builderType BuilderType

	for index, character := range []rune(line) {

		wasNumber := false

		if builder.Len() != 0 && (builderType == KeywordBuilderType || builderType == SymbolBuilderType) {
			// If matching token type is found, store token
			if tokenType, exists := l.syntax.tokenTypes[builder.String()]; exists {

				tokens = append(tokens, Token{
					ColumnRange: IntRange{Start: index - builder.Len(), End: index},
					LineNumber:  lineNumber,
					Type:        tokenType,
				})

				builder.Reset()
				builderType = UnsetBuilderType
			} else if builderType == KeywordBuilderType {

				if unicode.IsLetter(character) {
					builder.WriteRune(character)
					continue
				}

				tokens = append(tokens, Token{
					Value:       builder.String(),
					ColumnRange: IntRange{Start: index - builder.Len(), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.IdentifierTokenType,
				})

				builder.Reset()
				builderType = UnsetBuilderType
			}
		}

		switch character {

		case '"':
			if builderType == StringBuilderType {

				tokens = append(tokens, Token{
					Value:       builder.String(),
					ColumnRange: IntRange{Start: index - builder.Len() - len("\""), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.StringTokenType,
				})

				builder.Reset()
				builderType = UnsetBuilderType
			} else {
				builderType = StringBuilderType
			}

		case '\'':
			if builderType == CharBuilderType {

				tokens = append(tokens, Token{
					Value:       builder.String(),
					ColumnRange: IntRange{Start: index - builder.Len() - len("'"), End: index},
					LineNumber:  lineNumber,
					Type:        l.syntax.CharTokenType,
				})

				builder.Reset()
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
				builder.WriteRune(character)
				wasNumber = true
			}

		// Note: Negative numbers won't be lexed entirely, will be Minus token first
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

			// If there is a hanging dot, remove token and add to number
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == l.syntax.DotTokenType {
				tokens = tokens[:len(tokens)-1]
				builder.WriteRune('.')
			}

			builder.WriteRune(character)
			wasNumber = true

		default:

			switch builderType {

			case KeywordBuilderType:

			case SymbolBuilderType:

			case NumberBuilderType:

			}
			if unicode.IsLetter(character) {
				builder.WriteRune(character)
				builderType = KeywordBuilderType
			} else {
				builder.WriteRune(character)
				builderType = SymbolBuilderType
			}
		}

		if builderType == NumberBuilderType && !wasNumber && builder.Len() > 0 {
			// TODO: Store number as token
			// TODO: Account for a number with `f` as the suffix to be float
			builder.Reset()
			builderType = UnsetBuilderType
		}

	}

	// Must be identifier
	if builder.Len() > 0 {

		switch builderType {
		// TODO: Switch between types
		}

		tokens = append(tokens, Token{
			Value:       builder.String(),
			ColumnRange: IntRange{Start: len(line) - builder.Len() - 1, End: len(line) - 1},
			LineNumber:  lineNumber,
			Type:        l.syntax.IdentifierTokenType,
		})

		builder.Reset()
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
