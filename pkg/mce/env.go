package mce

import "github.com/TheDevtop/theta-go/pkg/sexp"

type Environment struct {
	parent *Environment
	table  map[sexp.Symbol]sexp.Value
}

func New() *Environment {
	ptr := new(Environment)
	ptr.parent = nil
	ptr.table = make(map[sexp.Symbol]sexp.Value, 8)
	return ptr
}

func (env *Environment) Lookup(sym sexp.Symbol) sexp.Value {
	if val, ok := env.table[sym]; ok {
		return val
	}
	if env.parent == nil {
		return nil
	}
	return env.parent.Lookup(sym)
}

func (env *Environment) Modify(sym sexp.Symbol, exp sexp.Value) {
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
