package types

// Environment
type Environment struct {
	parent *Environment
	table  map[Symbol]Value
}

// Parseable types
type (
	Symbol  string // + = foo bar
	Keyword string // :ok :err :name
	Value   any
	List    = []Value                      // ()
	Lambda  func(List, *Environment) Value // :fn
)
