package common

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
	Visibility Visibility
	Parameters []Parameter
}

type Parameter struct {
	Type Type
	Name string
}

type Variable struct {
	Value      Value
	Visibility *Visibility
	Name       string
	IsMutable  bool
}

type Operation struct {
	Node1     Node
	Node2     Node
	Operation Token
}

type Call struct {
	Next Node
	Name string
}

type ArrayType struct {
	Name       string
	Dimensions uint
}
