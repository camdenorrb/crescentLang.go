package common

import "github.com/moznion/go-optional"

// AST that can be used by any language for conversion purposes

type Node interface{}

type Value interface{}

// Type can be string or ArrayType
type Type interface{}

type Visibility string

type Parameters []Parameter

type Import struct {
	Value string
	Alias string
}

type Struct struct {
	Name      string
	Fields    Parameters
	Variables []Variable
	Functions []Function
}

type Enum struct {
}

type Function struct {
	Name       string
	Visibility Visibility
	Parameters []Parameter
	Body       []Node
	ReturnType Type
}

type Field struct {
	Type         Type
	Name         string
	DefaultValue optional.Option[Value]
}

type Parameter struct {
	Type         Type
	Name         string
	DefaultValue optional.Option[Value]
}

type Return struct {
	Value Value
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
	Name string
	Args []Node
	Next Node
}

type ArrayType struct {
	Name       string
	Dimensions uint
}
