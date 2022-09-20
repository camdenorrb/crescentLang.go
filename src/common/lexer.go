package common

import (
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
	syntax *Syntax
}

type tokenBuilder struct {
	syntax      *Syntax
	cache       strings.Builder
	tokens      []Token
	mode        BuilderMode
	columnIndex int
	lineNumber  uint
}

func (b *tokenBuilder) toggleMode(mode BuilderMode) error {

	if mode == UnsetBuilderMode {
		return nil
	}

	if b.mode != mode {

		if b.mode != UnsetBuilderMode {
			err := b.unsetMode()
			if err != nil {
				return errorx.Decorate(err, "[0] toggleMode failed to call unsetMode")
			}
		}

		b.mode = mode
		return nil
	}

	err := b.unsetMode()
	if err != nil {
		return errorx.Decorate(err, "[0] toggleMode failed to call unsetMode")
	}

	return nil
}

func (b *tokenBuilder) unsetMode() error {

	if b.mode == UnsetBuilderMode {
		return nil
	}

	token := Token{
		Value:       b.cache.String(),
		ColumnRange: IntRange{Start: b.columnIndex - b.cache.Len(), End: b.columnIndex},
		LineNumber:  b.lineNumber,
	}

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
			return errorx.IllegalState.New("Could not find matching symbol type for: %q", b.cache.String())
		}
		token.Type = tokenType

	case CommentBuilderMode:
		token.Type = b.syntax.CommentTokenType.Unwrap()

	case MultiLineCommentBuilderMode:
		token.Type = b.syntax.MultiLineCommentTokenType.Unwrap()

	}

	b.mode = UnsetBuilderMode
	b.tokens = append(b.tokens, token)
	b.cache.Reset()

	return nil
}

/*
Tries to find the symbol by searching for full length then truncating until matches
Then tries to repeat the process on the truncated data to find more matches and appends if so
Will return empty slice if none are found, will return string of no matches if only some are found
*/
func (b *tokenBuilder) findSymbols() ([]Token, string) {

	cacheAsString := b.cache.String()

	var tokens []Token

	// Early condition to avoid looping
	if tokenType, exists := b.syntax.tokenTypes[cacheAsString]; exists {

		tokens = append(tokens, Token{
			ColumnRange: IntRange{
				Start: b.columnIndex - b.cache.Len(),
				End:   b.columnIndex,
			},
			Type: tokenType,
		})

		return tokens, ""
	}

	start := 0
	current := cacheAsString

	for current != "" {
		for takeUntil := len(current); takeUntil > 0; takeUntil-- {

			if tokenType, exists := b.syntax.tokenTypes[cacheAsString]; exists {

				tokens = append(tokens, Token{
					ColumnRange: IntRange{
						Start: b.columnIndex - b.cache.Len() + start,
						End:   b.columnIndex - len(current),
					},
					Type: tokenType,
				})

				current = current[:]
			}

		}
	}

}

func (b *tokenBuilder) step(character rune) error {

	switch character {

	case '\n':
		b.lineNumber++
		b.columnIndex = 0
		return nil

	case '\'':
		err := b.toggleMode(CharBuilderMode)
		if err != nil {
			return errorx.Decorate(err, "[0] step failed to call toggleMode")
		}

	case '"':
		err := b.toggleMode(StringBuilderMode)
		if err != nil {
			return errorx.Decorate(err, "[1] step failed to call toggleMode")
		}

	// Note: Negative numbers won't be lexed entirely, will be Minus token first
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

		// If dot alone it should be originally be parsed as a symbol, but if number present, change
		if b.mode != UnsetBuilderMode && b.cache.String() != "." {
			err := b.unsetMode()
			if err != nil {
				return errorx.Decorate(err, "[0] step failed to call unsetMode")
			}
		}

		if b.mode != NumberBuilderMode {
			err := b.toggleMode(NumberBuilderMode)
			if err != nil {
				return errorx.Decorate(err, "[2] step failed to call toggleMode")
			}
		}

		b.cache.WriteRune(character)

	default:

		if b.mode != StringBuilderMode && b.mode != CharBuilderMode {

			// Fix mode
			switch {

			case unicode.IsSpace(character):
				err := b.unsetMode()
				if err != nil {
					return errorx.Decorate(err, "[1] step failed to call unsetMode")
				}
				b.columnIndex++
				return nil

			case unicode.IsLetter(character):
				if b.mode != IdentifierBuilderMode {
					err := b.toggleMode(IdentifierBuilderMode)
					if err != nil {
						return errorx.Decorate(err, "[3] step failed to call toggleMode")
					}
				}

			default:
				if b.mode != SymbolBuilderMode {
					err := b.toggleMode(SymbolBuilderMode)
					if err != nil {
						return errorx.Decorate(err, "[3] step failed to call toggleMode")
					}
				}
			}

		}

		b.cache.WriteRune(character)
	}

	b.columnIndex++

	return nil
}

func NewGenericLexer(syntax *Syntax) (*GenericLexer, error) {

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

	return &GenericLexer{syntax: syntax}, nil
}

func (l *GenericLexer) lex(reader *strings.Reader) ([]Token, error) {

	builder := tokenBuilder{syntax: l.syntax}

	for {

		character, _, err := reader.ReadRune()
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

	err := builder.unsetMode()
	if err != nil {
		return nil, errorx.Decorate(err, "lex failed to call unsetMode")
	}

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
