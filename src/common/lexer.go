package common

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"unicode"
)

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

type Lexer interface {
	Lex(lines chan string) []Token
}

type GenericLexer struct {
	reader *bufio.Reader
	syntax *Syntax
}

func NewGenericLexer(reader io.Reader, syntax *Syntax) *GenericLexer {
	return &GenericLexer{
		reader: bufio.NewReader(reader),
		syntax: syntax,
	}
}

// Plan: If syntax doesn't contain Rune, append and check again
// Make sure to check if number, string literal, etc

func (l *GenericLexer) lexLine(line string, lineNumber uint) ([]Token, error) {

	var tokens []Token

	var builder strings.Builder
	var builderType BuilderType

	for index, character := range []rune(line) {

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

		// Note: Negative numbers won't be lexed entirely, will be Minus token first
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

			// If dot alone it should be originally be parsed as a symbol, but if number present, change
			if builderType != UnsetBuilderType && builder.String() != "." {
				// Error
			}

			builder.WriteRune(character)
			builderType = NumberBuilderType

		default:

			switch builderType {

			case KeywordBuilderType:
				if !unicode.IsLetter(character) {
					// TODO: Error unexpected
					return nil, errors.New("")
				}
				builder.WriteRune(character)
			case SymbolBuilderType:
				builder.WriteRune(character)

			case NumberBuilderType:
				// If `f` store as float
				// Parse/Store number
				// If error parsing as number, error

				builder.Reset()
				builderType = UnsetBuilderType

			case UnsetBuilderType:
				if unicode.IsLetter(character) {
					builder.WriteRune(character)
					builderType = KeywordBuilderType
				} else {
					builder.WriteRune(character)
					builderType = SymbolBuilderType
				}
			}

		}
	}

	// Must be identifier
	if builder.Len() > 0 {

		switch builderType {
		// TODO: Switch between types, maybe have a function for handling each one that can be reused in default case
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
