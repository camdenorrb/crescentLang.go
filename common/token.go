package common

// Parser would essentially be a big switch statement in this case

type TokenType = uint8

// UIntRange is treated as inclusive
type UIntRange struct {
	Start uint
	End   uint
}

type Token struct {
	Value       interface{}
	ColumnRange UIntRange
	LineNumber  uint
	Type        TokenType
}
