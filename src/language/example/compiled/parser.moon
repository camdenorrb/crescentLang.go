# This will be compiled only if it's defined in the language file


fun ParseImport(import: String) -> Import {
	val import = import.split(" ")
	val name = import[1]
	val path = import[2]
	-> Import(name, path)
}