package types

/*
	Theta list processor
	Environment methods
*/

// Allocate and initialize environment with specified size
func NewEnvironment(size int) *Environment {
	var env = new(Environment)
	env.Parent = nil
	env.Table = make(map[Symbol]Expression, size)
	return env
}

// Recursive lookup from environment
func (env *Environment) Lookup(sym Symbol) Expression {
	if exp, ok := env.Table[sym]; ok {
		return exp
	}
	if env.Parent == nil {
		return nil
	}
	return env.Parent.Lookup(sym)
}

// Modify or delete from environment
func (env *Environment) Modify(sym Symbol, exp Expression) {
	if exp == nil {
		delete(env.Table, sym)
	}
	env.Table[sym] = exp
}

// Modify or delete from most parental environment
func (env *Environment) DeepModify(sym Symbol, exp Expression) {
	if env.Parent == nil {
		if exp == nil {
			delete(env.Table, sym)
		} else {
			env.Table[sym] = exp
		}
	} else {
		env.Parent.DeepModify(sym, exp)
	}
}

// Link parent from child environment
func (env *Environment) Link(penv *Environment) {
	if penv == nil || penv == env {
		return
	}
	env.Parent = penv
}

// Recursive dump of environment
func (env *Environment) Dump() List {
	var (
		sym  Symbol
		exp  Expression
		list = make(List, 0)
	)
	if env.Parent != nil {
		list = append(list, env.Parent.Dump()...)
	}
	for sym, exp = range env.Table {
		list = append(list, List{sym, exp})
	}
	return list
}
