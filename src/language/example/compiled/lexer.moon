import io

# This will be compiled only if it's defined in the language file

type TokenType = uint
const (
	IF TokenType = iota
)

val generatedTokens = map[String]TokenType {
	"if": TokenType.IF
}

fun generatedLex(reader: io.Reader) -> [Token] {
	reader.nw
}