package common

import (
	"fmt"
	"github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestGenericLexer_lexLine(t *testing.T) {
	type args struct {
		line       string
		lineNumber uint
	}
	tests := []struct {
		name    string
		args    args
		want    []Token
		wantErr bool
	}{
		{
			name: "Meow",
			args: args{
				line: "println;(\"Meow\");",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l, err := NewGenericLexer(&Syntax{
				KeyTokenTypes: KeyTokenTypes{
					IdentifierTokenType:       optional.Some(TokenType(2)),
					CharTokenType:             optional.Some(TokenType(3)),
					StringTokenType:           optional.Some(TokenType(4)),
					NumberTokenType:           optional.Some(TokenType(5)),
					CommentTokenType:          optional.Some(TokenType(6)),
					MultiLineCommentTokenType: optional.Some(TokenType(7)),
				},
				tokenTypes: map[string]TokenType{
					"(": TokenType(0),
					")": TokenType(1),
				},
			})
			assert.NoError(t, err)
			assert.NotNil(t, l)

			got, err := l.lex(strings.NewReader(tt.args.line))

			for _, token := range got {
				fmt.Printf("%+v\n", token)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("lexLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexLine() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestTanna(t *testing.T) {

	lexer, err := NewGenericLexer(&Syntax{
		KeyTokenTypes: KeyTokenTypes{
			IdentifierTokenType:       optional.Some(TokenType(2)),
			CharTokenType:             optional.Some(TokenType(3)),
			StringTokenType:           optional.Some(TokenType(4)),
			NumberTokenType:           optional.Some(TokenType(5)),
			CommentTokenType:          optional.Some(TokenType(6)),
			MultiLineCommentTokenType: optional.Some(TokenType(7)),
		},
		tokenTypes: map[string]TokenType{
			"sout":     TokenType(0),
			"function": TokenType(1),
			"Int":      TokenType(2),
			"(":        TokenType(3),
			")":        TokenType(4),
			"+":        TokenType(5),
			"=>":       TokenType(6),
			"}":        TokenType(7),
			"{":        TokenType(8),
			":":        TokenType(9),
			",":        TokenType(10),
			"=":        TokenType(11),
			">":        TokenType(12),
		},
	})
	assert.NoError(t, err)

	input := `
			constant value0 11
			constant value1 = 20

			function add(a: Int, b: Int): Int {
				=> a + b
			}
		`

	tokens, err := lexer.lex(strings.NewReader(input))
	assert.NoError(t, err)

	lines := strings.Split(input, "\n")

	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}

	for _, token := range tokens {
		fmt.Println(lines[token.LineNumber][token.ColumnRange.Start:token.ColumnRange.End], len(lines[token.LineNumber][token.ColumnRange.Start:token.ColumnRange.End]))
	}
}

func BenchmarkWord(b *testing.B) {

	lexer, err := NewGenericLexer(&Syntax{
		KeyTokenTypes: KeyTokenTypes{
			IdentifierTokenType:       optional.Some(TokenType(2)),
			CharTokenType:             optional.Some(TokenType(3)),
			StringTokenType:           optional.Some(TokenType(4)),
			NumberTokenType:           optional.Some(TokenType(5)),
			CommentTokenType:          optional.Some(TokenType(6)),
			MultiLineCommentTokenType: optional.Some(TokenType(7)),
		},
		tokenTypes: map[string]TokenType{
			"sout":     TokenType(0),
			"function": TokenType(1),
			"Int":      TokenType(2),
			"(":        TokenType(3),
			")":        TokenType(4),
			"+":        TokenType(5),
			"=>":       TokenType(6),
			"}":        TokenType(7),
			"{":        TokenType(8),
			":":        TokenType(9),
			",":        TokenType(10),
		},
	})
	if err != nil {
		return
	}

	reader := strings.NewReader(`sout "Hello World!"`)

	for i := 0; i < b.N; i++ {
		_, err = lexer.lex(reader)
		if err != nil {
			b.Error(err)
		}
		reader.Reset(`sout "Hello World!"`)
	}
}
