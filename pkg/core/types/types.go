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
	Set        map[Symbol]bool
	Function   struct {
		Args Set
		Body Expression
	}
	Procedure func(*Environment, ...Expression) Expression
)
