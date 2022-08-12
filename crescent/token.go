package crescent

import (
	"crescentLang/global"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

const (
	NOT global.TokenType = iota
	EQUALS
	SEMICOLON
	TYPE_PREFIX

	// ()
	PARENL
	PARENR

	// {}
	BRACEL
	BRACER

	// []
	BRACKETL
	BRACKETR

	// Infix Operators
	CONTAINS
	RANGE_TO
	AS

	// Variables
	VAR
	VAL
	CONST

	// Types
	STRUCT
	IMPL
	TRAIT
	OBJECT
	ENUM
	SEALED

	// Statements
	ELSE
	IMPORT
	IF
	WHEN
	WHILE
	FOR
	FUN

	// Modifiers
	ASYNC
	OVERRIDE
	OPERATOR
	INLINE
	STATIC

	// Visibility
	PUBLIC
	INTERNAL
	PRIVATE

	// Arithmetic
	ADD
	SUB
	MUL
	DIV
	REM
	POW

	// Bit
	SHL
	SHR
	USHR
	AND
	OR
	XOR

	// Assign
	ASSIGN
	ADD_ASSIGN
	SUB_ASSIGN
	MUL_ASSIGN
	DIV_ASSIGN
	REM_ASSIGN
	POW_ASSIGN

	// Compare
	OR_COMPARE
	AND_COMPARE
	LESSER_THAN_COMPARE
	GREATER_THAN_COMPARE
	LESSER_THAN_OR_EQUALS_COMPARE
	GREATER_THAN_OR_EQUALS_COMPARE
	EQUALS_COMPARE
	NOT_EQUALS_COMPARE

	// Compare references
	EQUALS_REFERENCE_COMPARE
	NOT_EQUALS_REFERENCE_COMPARE

	// Bool
	TRUE
	FALSE

	INSTANCE_OF
	RETURN
	RESULT
	COMMA
	DOT
	IMPORT_SEPARATOR
)

var tokenTypes = linkedhashset.New(

	"!",
	"=",
	";",
	":",

	"(",
	")",

	"{",
	"}",

	"[",
	"]",

	"in",
	"..",
	"as",

	"var",
	"val",
	"const",

	"struct",
	"impl",
	"trait",
	"object",
	"enum",
	"sealed",

	"else",
	"import",
	"if",
	"when",
	"while",
	"for",
	"fun",

	"async",
	"override",
	"operator",
	"inline",
	"static",

	"public",
	"internal",
	"private",

	"+",
	"-",
	"*",
	"/",
	"%",
	"^",

	"shl",
	"shr",
	"ushr",
	"and",
	"or",
	"xor",

	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"^=",

	"|",
	"&",
	"<",
	">",
	"<=",
	">=",
	"==",
	"!=",

	"===",
	"!==",

	"true",
	"false",

	"is",
	"->",
	"?",
	",",
	".",
	"::",
)
