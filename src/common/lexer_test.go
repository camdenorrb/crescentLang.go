package common

import (
	"github.com/markphelps/optional"
	"reflect"
	"testing"
	"unicode"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l := NewGenericLexer(nil, &Syntax{
				identifierTokenType:          0,
				identifierCharacterValidator: unicode.IsLetter,
				tokenTypes:                   nil,
				commentPrefix:                optional.String{},
				stringWrapper:                0,
				charWrapper:                  0,
			})
			got, err := l.lexLine(tt.args.line, tt.args.lineNumber)
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
