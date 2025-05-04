package mce

import (
	"github.com/TheDevtop/theta-go/pkg/sexp"
)

var (
	ErrInvalidArgs = sexp.Message(sexp.KeyErr, "Invalid arguments")
	errInvalidEval = sexp.Message(sexp.KeyErr, "Invalid evaluation form")
)

// (fn (ARGS) BODY)
func applyFn(cdr sexp.List, env *Environment) sexp.Value {
	var (
		ok      bool
		argList sexp.List
		body    sexp.Value
		fn      Lambda
	)

	// Check arguments
	if len(cdr) != 2 {
		return ErrInvalidArgs
	}
	if argList, ok = cdr[0].(sexp.List); !ok {
		return ErrInvalidArgs
	}
	if !sexp.IsConsistent[sexp.Symbol](argList) {
		return ErrInvalidArgs
	}
	body = cdr[1]

	// Construct function
	fn = func(args sexp.List, penv *Environment) sexp.Value {
		var (
			env = New()
		)
		if len(args) != len(argList) {
			return ErrInvalidArgs
		}
		for i, sym := range argList {
			env.Modify(sym.(sexp.Symbol), args[i])
		}
		env.Link(penv)
		return Eval(body, env)
	}
	return fn
}

// (def SYMBOL VALUE)
func applyDef(cdr sexp.List, env *Environment) sexp.Value {
	var (
		sym sexp.Symbol
		ok  bool
	)
	if len(cdr) != 2 {
		return ErrInvalidArgs
	}
	if sym, ok = cdr[0].(sexp.Symbol); !ok {
		return ErrInvalidArgs
	}
	env.Modify(sym, Eval(cdr[1], env))
	return sexp.KeyOk
}

// (if COND THEN ELSE)
func applyIf(cdr sexp.List, env *Environment) sexp.Value {
	var (
		condval bool
		ok      bool
	)
	if len(cdr) != 3 {
		return ErrInvalidArgs
	}
	if condval, ok = Eval(cdr[0], env).(bool); !ok {
		return ErrInvalidArgs
	}
	if condval {
		return Eval(cdr[1], env)
	}
	return Eval(cdr[2], env)
}

func Apply(fn Lambda, args sexp.List, env *Environment) sexp.Value {
	nargs := make(sexp.List, len(args))
	for i, arg := range args {
		nargs[i] = Eval(arg, env)
	}
	return fn(nargs, env)
}

func evalCons(car sexp.Value, cdr sexp.List, env *Environment) sexp.Value {
	var (
		op sexp.Symbol
		ok bool
		fn Lambda
	)
	if op, ok = car.(sexp.Symbol); !ok {
		return errInvalidEval
	}
	switch op {
	case "quote":
		if len(cdr) != 1 {
			return ErrInvalidArgs
		}
		return cdr[0]
	case "if":
		return applyIf(cdr, env)
	case "def":
		return applyDef(cdr, env)
	case "fn":
		return applyFn(cdr, env)
	}
	if fn, ok = env.Lookup(op).(Lambda); !ok {
		return errInvalidEval
	}
	return Apply(fn, cdr, env)
}

func Eval(exp sexp.Value, env *Environment) sexp.Value {
	switch exp.(type) {
	case sexp.Symbol:
		return env.Lookup(exp.(sexp.Symbol))
	case sexp.List:
		car, cdr := sexp.Cons(exp.(sexp.List))
		return evalCons(car, cdr, env)
	default:
		return exp
	}
}
