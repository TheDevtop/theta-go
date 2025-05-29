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
	Symbol  string
	Keyword string
	Number  struct {
		int
		float64
	}
	Expression any
	List       = []Expression
	Function   func(*Environment, ...Expression) Expression
)
