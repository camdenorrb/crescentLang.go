package global

// AST that can be used by any language for conversion purposes

type Node interface{}

type Value interface{}

// Type can be string or ArrayType
type Type interface{}

type Visibility string

type Struct struct {
	Name      string
	Parameter []Parameter
}

type Function struct {
	Name       string
	Parameters []Parameter
	Visibility Visibility
}

type Parameter struct {
	Name string
	Type Type
}

type Variable struct {
	Name       string
	Value      Value
	IsMutable  bool
	Visibility *Visibility
}

type Operation struct {
	Node1     Node
	Node2     Node
	Operation Token
}

type Call struct {
	Name string
	Next Node
}

type ArrayType struct {
	Name       string
	Dimensions uint
}
