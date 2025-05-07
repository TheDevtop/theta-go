package types

/*
	Theta list processor
	Core types
*/

type Environment struct {
	Parent *Environment
	Table  map[Symbol]Expression
}

type (
	Symbol     string
	Keyword    string
	Expression any
	List       = []Expression
	Function   func(*Environment, ...Expression) Expression
)
