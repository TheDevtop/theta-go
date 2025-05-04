package types

func NewEnvironment() *Environment {
	ptr := new(Environment)
	ptr.parent = nil
	ptr.table = make(map[Symbol]Value, 8)
	return ptr
}

func (env *Environment) Lookup(sym Symbol) Value {
	if val, ok := env.table[sym]; ok {
		return val
	}
	if env.parent == nil {
		return nil
	}
	return env.parent.Lookup(sym)
}

func (env *Environment) Modify(sym Symbol, exp Value) {
	if exp == nil {
		delete(env.table, sym)
	}
	env.table[sym] = exp
}

func (env *Environment) Link(penv *Environment) {
	if penv == nil || penv == env {
		return
	}
	env.parent = penv
}
