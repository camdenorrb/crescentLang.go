package common

import (
	"bufio"
	"errors"
	"github.com/joomcode/errorx"
	"io"
	"strings"
	"unicode"
)

type BuilderMode uint

const (
	UnsetBuilderMode BuilderMode = iota
	IdentifierBuilderMode
	SymbolBuilderMode
	StringBuilderMode
	CharBuilderMode
	NumberBuilderMode
	CommentBuilderMode
	MultiLineCommentBuilderMode
)

type Lexer interface {
	Lex(lines chan string) []Token
}

type GenericLexer struct {
	reader      *bufio.Reader
	syntax      *Syntax
	builderMode BuilderMode
	tokens      []Token
}

type tokenBuilder struct {
	mode        BuilderMode
	cache       strings.Builder
	tokens      []Token
	columnIndex int
	lineNumber  uint
	syntax      *Syntax
}

func (b *tokenBuilder) toggleMode(mode BuilderMode) {

	if mode == UnsetBuilderMode {
		return
	}

	if b.mode != mode {

		if b.mode != UnsetBuilderMode {
			b.unsetMode()
		}

		b.mode = mode
		return
	}

	b.unsetMode()
}

func (b *tokenBuilder) unsetMode() {

	if b.mode == UnsetBuilderMode {
		return
	}

	token := Token{
		Value:       b.cache.String(),
		ColumnRange: IntRange{Start: b.columnIndex - b.cache.Len(), End: b.columnIndex},
		LineNumber:  b.lineNumber,
	}

	b.cache.Reset()

	switch b.mode {

	case CharBuilderMode:
		// TODO: Validate length of cache
		token.ColumnRange.Start -= 1 // To account for starting '
		token.Type = b.syntax.CharTokenType.Unwrap()

	case StringBuilderMode:
		token.ColumnRange.Start -= 1 // To account for starting "
		token.Type = b.syntax.StringTokenType.Unwrap()

	case NumberBuilderMode:
		// TODO: Validate
		token.Type = b.syntax.NumberTokenType.Unwrap()

	case IdentifierBuilderMode:
		token.Type = b.syntax.IdentifierTokenType.Unwrap()

	case SymbolBuilderMode:
		tokenType, exists := b.syntax.tokenTypes[b.cache.String()]
		if !exists {
			// TODO: Error
		}
		token.Type = tokenType

	case CommentBuilderMode:
		token.Type = b.syntax.CommentTokenType.Unwrap()

	case MultiLineCommentBuilderMode:
		token.Type = b.syntax.MultiLineCommentTokenType.Unwrap()

	}

	b.mode = UnsetBuilderMode
	b.tokens = append(b.tokens, token)
}

func (b *tokenBuilder) step(character rune) error {

	switch character {

	case '\n':
		b.lineNumber++
		b.columnIndex = 0
		return nil

	case '\'':
		b.toggleMode(CharBuilderMode)

	case '"':
		b.toggleMode(StringBuilderMode)

	// Note: Negative numbers won't be lexed entirely, will be Minus token first
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

		// If dot alone it should be originally be parsed as a symbol, but if number present, change
		if b.mode != UnsetBuilderMode && b.cache.String() != "." {
			return errorx.IllegalState.New("invalid number builder %q lineNumber: %d, column: %d", b.cache.String(), b.lineNumber, b.columnIndex)
		}

		b.cache.WriteRune(character)

		if b.mode == UnsetBuilderMode {
			b.toggleMode(NumberBuilderMode)
		}

	default:

		if b.mode == NumberBuilderMode {
			b.unsetMode()
		}

		if b.mode != StringBuilderMode && b.mode != CharBuilderMode {

			isLetter := unicode.IsLetter(character)

			if b.mode != IdentifierBuilderMode && isLetter {
				b.toggleMode(IdentifierBuilderMode)
			} else if b.mode != SymbolBuilderMode && !isLetter {
				b.toggleMode(SymbolBuilderMode)
			}
		}

		b.cache.WriteRune(character)
	}

	b.columnIndex++

	return nil
}

func NewGenericLexer(reader io.Reader, syntax *Syntax) (*GenericLexer, error) {

	if syntax.CharTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires CharTokenType not to be nil")
	}

	if syntax.StringTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires StringTokenType not to be nil")
	}

	if syntax.NumberTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires NumberTokenType not to be nil")
	}

	if syntax.IdentifierTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires IdentifierTokenType not to be nil")
	}

	if syntax.CommentTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires CommentTokenType not to be nil")
	}

	if syntax.MultiLineCommentTokenType.IsNone() {
		return nil, errors.New("NewGenericLexer requires MultiLineCommentTokenType not to be nil")
	}

	return &GenericLexer{
		reader:      bufio.NewReader(reader),
		syntax:      syntax,
		builderMode: UnsetBuilderMode,
	}, nil
}

func (l *GenericLexer) lex() ([]Token, error) {

	builder := tokenBuilder{syntax: l.syntax}

	for {

		character, _, err := l.reader.ReadRune()
		if err != nil {

			if err == io.EOF {
				break
			}

			return nil, errorx.Decorate(err, "lex failed to read rune")
		}

		err = builder.step(character)
		if err != nil {
			return nil, errorx.Decorate(err, "lex failed to step")
		}
	}

	builder.unsetMode()

	return builder.tokens, nil
}

/*
func (l *GenericLexer) lex() error {

	var builder strings.Builder

	for {

		character, _, err := l.reader.ReadRune()
		if err != nil {

			if err == io.EOF {
				break
			}

			return errorx.Decorate(err, "lex failed to read rune")
		}

		columnIndex++

		switch character {

		case '\n':
			lineNumber++
			columnIndex = 0

		case '\'':
			l.builderMode = CharBuilderMode

		case '"':
			l.builderMode = StringBuilderMode

			builder.WriteRune(character)
		}

	}

	// Plan: If syntax doesn't contain Rune, append and check again
	// Make sure to check if number, string literal, etc

	// TODO: On new line stop comment builder type
	// TODO: Make a switch for builder type that will automagically save token for last type
	func(l *GenericLexer) lexLine(line
	string, lineNumber
	uint) ([]Token, error) {

		var tokens []Token

		var builder strings.Builder

		for index, character := range []rune(line) {

			if builder.Len() != 0 && (builderMode == IdentifierBuilderMode || builderMode == SymbolBuilderMode) {
				// If matching token type is found, store token
				if tokenType, exists := l.syntax.tokenTypes[builder.String()]; exists {

					tokens = append(tokens, Token{
						ColumnRange: IntRange{Start: index - builder.Len(), End: index},
						LineNumber:  lineNumber,
						Type:        tokenType,
					})

					builder.Reset()
					builderMode = UnsetBuilderMode
				} else if builderMode == IdentifierBuilderMode {

					if unicode.IsLetter(character) {
						builder.WriteRune(character)
						continue
					}

					tokens = append(tokens, Token{
						Value:       builder.String(),
						ColumnRange: IntRange{Start: index - builder.Len(), End: index},
						LineNumber:  lineNumber,
						Type:        l.syntax.IdentifierTokenType.Unwrap(),
					})

					builder.Reset()
					builderMode = UnsetBuilderMode
				}
			}

			switch character {

			case '"':
				if builderMode == StringBuilderMode {

					tokens = append(tokens, Token{
						Value:       builder.String(),
						ColumnRange: IntRange{Start: index - builder.Len() - len("\""), End: index},
						LineNumber:  lineNumber,
						Type:        l.syntax.StringTokenType.Unwrap(),
					})

					builder.Reset()
					builderMode = UnsetBuilderMode
				} else {
					builderMode = StringBuilderMode
				}

			case '\'':
				if builderMode == CharBuilderMode {

					tokens = append(tokens, Token{
						Value:       builder.String(),
						ColumnRange: IntRange{Start: index - builder.Len() - len("'"), End: index},
						LineNumber:  lineNumber,
						Type:        l.syntax.CharTokenType.Unwrap(),
					})

					builder.Reset()
					builderMode = UnsetBuilderMode
				} else {
					builderMode = CharBuilderMode
				}

			// Note: Negative numbers won't be lexed entirely, will be Minus token first
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

				// If dot alone it should be originally be parsed as a symbol, but if number present, change
				if builderMode != UnsetBuilderMode && builder.String() != "." {
					return nil, errorx.IllegalState.New("invalid number builder %q lineNumber: %d, column: %d", builder.String(), lineNumber, index)
				}

				builder.WriteRune(character)
				builderMode = NumberBuilderMode

			default:

				switch builderMode {

				case IdentifierBuilderMode:
					if !unicode.IsLetter(character) {
						return nil, errorx.IllegalState.New("invalid keyword character %q lineNumber: %d, column: %d", character, lineNumber, index)
					}
					builder.WriteRune(character)

				case SymbolBuilderMode:
					if unicode.IsLetter(character) || unicode.IsDigit(character) {
						return nil, errorx.IllegalState.New("invalid symbol character %q lineNumber: %d, column: %d", character, lineNumber, index)
					}
					builder.WriteRune(character)

				case StringBuilderMode:
					builder.WriteRune(character)

				case UnsetBuilderMode:
					if unicode.IsLetter(character) {
						builder.WriteRune(character)
						builderMode = IdentifierBuilderMode
					} else {
						builder.WriteRune(character)
						builderMode = SymbolBuilderMode
					}

				case NumberBuilderMode:

					if unicode.ToLower(character) != 'f' {
						return nil, errorx.IllegalState.New("invalid number character %q lineNumber: %d, column: %d", character, lineNumber, index)
					}

					asFloat, err := strconv.ParseFloat(builder.String(), 32)
					if err != nil {
						return nil, errorx.IllegalState.New("invalid number builder %q lineNumber: %d, column: %d", builder.String(), lineNumber, index)
					}

					tokens = append(tokens, Token{
						Value:       float32(asFloat),
						ColumnRange: IntRange{Start: builder.Len() - index, End: index},
						LineNumber:  lineNumber,
						Type:        l.syntax.NumberTokenType.Unwrap(),
					})

					builder.Reset()
					builderMode = UnsetBuilderMode
				}
			}
		}

		// Must be identifier
		if builder.Len() > 0 {

			builderAsString := builder.String()

			if tokenType, ok := l.syntax.tokenTypes[builderAsString]; ok {

				tokens = append(tokens, Token{
					ColumnRange: IntRange{Start: len(line) - builder.Len() - 1, End: len(line) - 1},
					LineNumber:  lineNumber,
					Type:        tokenType,
				})

			} else {
				switch builderMode {

				// TODO: Switch between types, maybe have a function for handling each one that can be reused in default case
				case NumberBuilderMode:

					asFloat, err := strconv.ParseFloat(builderAsString, 64)
					if err != nil {
						return nil, errorx.IllegalState.New("invalid number builder %q lineNumber: %d, column: %d", builder.String(), lineNumber, len(line)-builder.Len()-1)
					}

					tokens = append(tokens, Token{
						Value:       asFloat,
						ColumnRange: IntRange{Start: len(line) - builder.Len() - 1, End: len(line) - 1},
						LineNumber:  lineNumber,
						Type:        l.syntax.NumberTokenType.Unwrap(),
					})

				case IdentifierBuilderMode:
					tokens = append(tokens, Token{
						Value:       builderAsString,
						ColumnRange: IntRange{Start: len(line) - builder.Len() - 1, End: len(line) - 1},
						LineNumber:  lineNumber,
						Type:        l.syntax.IdentifierTokenType.Unwrap(),
					})

				default:
					return nil, errorx.IllegalState.New("unexpected builder type %q for %q", builderMode, builderAsString)

				}
			}

			builder.Reset()
			builderMode = UnsetBuilderMode
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
