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
				line: "println(\"Meow\")",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l, err := NewGenericLexer(strings.NewReader(tt.args.line), &Syntax{
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

			got, err := l.lex()

			fmt.Println(tt.args.line[14:15])
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
