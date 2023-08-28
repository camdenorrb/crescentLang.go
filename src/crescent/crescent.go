package crescent

import (
	"crescentLang/common"
	"unicode"
)

func Lex(lines chan string) []common.Token {

	var tokens []common.Token

	lineNumber := uint(1)

	for line := range lines {
		lineTokens := lexLine(line, lineNumber)
		tokens = append(tokens, lineTokens...)
		lineNumber++
	}

	return tokens
}

func lexLine(line string, lineNumber uint) []common.Token {

	var tokens []common.Token
	var number string
	var identifier string

	for index, character := range []rune(line) {

		wasNumber := false

		if identifier != "" {

			if unicode.IsLetter(character) {
				identifier += string(character)
				continue
			}

			// TODO: Check for keywords
			// TODO: This is a very good way to abstract this, keywords can be defined by the language Map<String, TokenType>, the rest should be the same for any language

			tokens = append(tokens, common.Token{
				Value:       identifier,
				ColumnRange: common.IntRange{Start: index - len(identifier), End: index},
				LineNumber:  lineNumber,
				Type:        IDENTIFIER,
			})

			identifier = ""
		}

		if unicode.IsSpace(character) {
			continue
		}

		switch character {

		case '+':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        ADD,
			})

		case '-':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        SUB,
			})

		case '*':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        MUL,
			})

		case '/':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        DIV,
			})

		case '%':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        REM,
			})

		case '^':
			tokens = append(tokens, common.Token{
				ColumnRange: common.IntRange{Start: index, End: index},
				LineNumber:  lineNumber,
				Type:        POW,
			})

		case '=':
			// TODO: Peek last token, pop if add, sub, etc

		case '.':
			if len(number) > 0 {
				number += "."
				wasNumber = true
			} else {
				tokens = append(tokens, common.Token{
					ColumnRange: common.IntRange{Start: index, End: index},
					LineNumber:  lineNumber,
					Type:        DOT,
				})
			}

		// Note: Negative numbers won't be lexed entirely, will be Minus symbol first
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':

			// If there is a hanging dot, remove token and add to number
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == DOT {
				tokens = tokens[:len(tokens)-1]
				number += "."
			}

			number += string(character)
			wasNumber = true

		default:

			if unicode.IsLetter(character) {
				identifier += string(character)
			}

			// TODO: Error token not found
		}

		if !wasNumber && number != "" {
			// TODO: Store number as token
			number = ""
		}

	}

	// TODO: Store number as token if not stored

}

func Parse() {

}
