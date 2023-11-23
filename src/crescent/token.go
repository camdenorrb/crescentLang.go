package crescent

import (
	"crescentLang/common"
	"github.com/moznion/go-optional"
)

// TODO: Move this to a function that returns a Syntax object
var Syntax = &common.Syntax{
	KeyTokenTypes: common.KeyTokenTypes{
		IdentifierTokenType:       optional.Some(IDENTIFIER),
		CharTokenType:             optional.Some(CHAR),
		StringTokenType:           optional.Some(STRING),
		NumberTokenType:           optional.Some(NUMBER),
		CommentTokenType:          optional.Some(COMMENT),
		MultiLineCommentTokenType: optional.Some(COMMENT),
	},
	CommentPrefix: optional.Some("#"),
	TokenTypes:    tokens,
}

var tokens = map[string]common.TokenType{
	"!": NOT,
	";": SEMICOLON,
	":": TYPE_PREFIX,

	// ()
	"(": PARENL,
	")": PARENR,

	// {}
	"{": CURLYL,
	"}": CURLYR,

	// []
	"[": BRACKETL,
	"]": BRACKETR,

	// Infix Operators
	"in": CONTAINS,
	"..": RANGE_TO,
	"as": AS,

	// Variables
	"var":   VAR,
	"val":   VAL,
	"const": CONST,

	// Types
	"struct": STRUCT,
	"impl":   IMPL,
	"trait":  TRAIT,
	"object": OBJECT,
	"enum":   ENUM,
	"sealed": SEALED,

	// Statements
	"else":   ELSE,
	"import": IMPORT,
	"if":     IF,
	"when":   WHEN,
	"while":  WHILE,
	"for":    FOR,
	"fun":    FUN,

	// Modifiers
	"async":    ASYNC,
	"override": OVERRIDE,
	"operator": OPERATOR,
	"inline":   INLINE,
	"static":   STATIC,

	// Visibility
	"public":   PUBLIC,
	"internal": INTERNAL,
	"private":  PRIVATE,

	// Arithmetic
	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
	"%": REM,
	"^": POW,

	// Bit
	"shl":  SHL,
	"shr":  SHR,
	"ushr": USHR,
	"and":  AND,
	"or":   OR,
	"xor":  XOR,

	// Assign
	"=":  ASSIGN,
	"+=": ADD_ASSIGN,
	"-=": SUB_ASSIGN,
	"*=": MUL_ASSIGN,
	"/=": DIV_ASSIGN,
	"%=": REM_ASSIGN,
	"^=": POW_ASSIGN,

	// Compare
	"|":  OR_COMPARE,
	"&":  AND_COMPARE,
	"<":  LESSER_THAN_COMPARE,
	">":  GREATER_THAN_COMPARE,
	"<=": LESSER_THAN_OR_EQUALS_COMPARE,
	">=": GREATER_THAN_OR_EQUALS_COMPARE,
	"==": EQUALS_COMPARE,
	"!=": NOT_EQUALS_COMPARE,

	// Compare references
	"===": EQUALS_REFERENCE_COMPARE,
	"!==": NOT_EQUALS_REFERENCE_COMPARE,

	// Bool
	"true":  TRUE,
	"false": FALSE,

	"is": INSTANCE_OF,
	"->": RETURN,
	"?":  RESULT,
	",":  COMMA,
	".":  DOT,
	"::": IMPORT_SEPARATOR,
}

const (
	NOT common.TokenType = iota
	SEMICOLON
	TYPE_PREFIX

	// ()
	PARENL
	PARENR

	// {}
	CURLYL
	CURLYR

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
	IDENTIFIER

	// Values
	STRING
	CHAR
	NUMBER
	COMMENT
)

/*
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
}
*/
