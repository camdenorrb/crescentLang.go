package common

// Parser would essentially be a big switch statement in this case

type TokenType = uint8

// IntRange is treated as inclusive
type IntRange struct {
	Start int
	End   int
}

type Token struct {
	Value       any
	ColumnRange IntRange
	LineNumber  uint
	Type        TokenType
}
