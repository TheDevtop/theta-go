package mce

import (
	"github.com/TheDevtop/theta-go/pkg/types"
)

var (
	ErrInvalidArgs = types.Message(types.KeyErr, "Invalid arguments")
	errInvalidEval = types.Message(types.KeyErr, "Invalid evaluation form")
)

// (fn (ARGS) BODY)
func applyFn(cdr types.List, env *types.Environment) types.Value {
	var (
		ok      bool
		argList types.List
		body    types.Value
		fn      types.Lambda
	)

	// Check arguments
	if len(cdr) != 2 {
		return ErrInvalidArgs
	}
	if argList, ok = cdr[0].(types.List); !ok {
		return ErrInvalidArgs
	}
	if !types.IsConsistent[types.Symbol](argList) {
		return ErrInvalidArgs
	}
	body = cdr[1]

	// Construct function
	fn = func(args types.List, penv *types.Environment) types.Value {
		var (
			env = types.NewEnvironment()
		)
		if len(args) != len(argList) {
			return ErrInvalidArgs
		}
		for i, sym := range argList {
			env.Modify(sym.(types.Symbol), args[i])
		}
		env.Link(penv)
		return Eval(body, env)
	}
	return fn
}

// (def SYMBOL VALUE)
func applyDef(cdr types.List, env *types.Environment) types.Value {
	var (
		sym types.Symbol
		ok  bool
	)
	if len(cdr) != 2 {
		return ErrInvalidArgs
	}
	if sym, ok = cdr[0].(types.Symbol); !ok {
		return ErrInvalidArgs
	}
	env.Modify(sym, Eval(cdr[1], env))
	return types.KeyOk
}

// (if COND THEN ELSE)
func applyIf(cdr types.List, env *types.Environment) types.Value {
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

// (seq SEQ...)
func applySeq(cdr types.List, env *types.Environment) types.Value {
	var val types.Value
	for _, exp := range cdr {
		val = Eval(exp, env)
	}
	return val
}

func Apply(fn types.Lambda, args types.List, env *types.Environment) types.Value {
	nargs := make(types.List, len(args))
	for i, arg := range args {
		nargs[i] = Eval(arg, env)
	}
	return fn(nargs, env)
}

func evalCons(car types.Value, cdr types.List, env *types.Environment) types.Value {
	var (
		op types.Symbol
		ok bool
		fn types.Lambda
	)
	if op, ok = car.(types.Symbol); !ok {
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
	case "seq":
		return applySeq(cdr, env)
	}
	if fn, ok = env.Lookup(op).(types.Lambda); !ok {
		return errInvalidEval
	}
	return Apply(fn, cdr, env)
}

func Eval(exp types.Value, env *types.Environment) types.Value {
	switch exp.(type) {
	case types.Symbol:
		return env.Lookup(exp.(types.Symbol))
	case types.List:
		car, cdr := types.Cons(exp.(types.List))
		return evalCons(car, cdr, env)
	default:
		return exp
	}
}
