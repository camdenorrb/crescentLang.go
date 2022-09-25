package common

// Parser would essentially be a big switch statement in this case

type TokenType = uint8

// IntRange is treated as inclusive
type IntRange struct {
	Start int
	End   int
}

type Token struct {
	//LineRange   optional.Option[IntRange] Not needed just parse multiline comments as multiple comments
	Value       any
	ColumnRange IntRange
	LineNumber  uint
	Type        TokenType
}
