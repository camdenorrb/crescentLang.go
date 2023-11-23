
# There should be a program that takes in all these files and creates a CLI for it
# Can be used for all language definitions
# The CLI can also have the conversion functionality from [Node]

fun Lex() -> [Token] {
	-> generatedLex()
}

fun Parse(tokens: [Token]) -> [Node] {
	-> generatedParse()
}