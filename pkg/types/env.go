package types

/*
	Theta virtual machine
	Environment functions
*/

// Allocate new environment
func NewEnvironment() *Environment {
	ptr := new(Environment)
	ptr.parent = nil
	ptr.table = make(map[Symbol]Value, 8)
	return ptr
}

// Initialize environment from existing table
func InitEnvironment(tab map[Symbol]Value) *Environment {
	ptr := new(Environment)
	ptr.parent = nil
	ptr.table = tab
	return ptr
}

// Recursive lookup from environment
func (env *Environment) Lookup(sym Symbol) Value {
	if val, ok := env.table[sym]; ok {
		return val
	}
	if env.parent == nil {
		return nil
	}
	return env.parent.Lookup(sym)
}

// Modify or delete from environment
func (env *Environment) Modify(sym Symbol, exp Value) {
	if exp == nil {
		delete(env.table, sym)
	}
	env.table[sym] = exp
}

// Link parent from child environment
func (env *Environment) Link(penv *Environment) {
	if penv == nil || penv == env {
		return
	}
	env.parent = penv
}
