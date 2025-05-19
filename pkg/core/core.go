package core

/*
	Theta list processor
	Evaluator
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

// Construct and apply arguments to function
func apply(env *types.Environment, fn types.Function, exp ...types.Expression) types.Expression {
	var args types.List = make(types.List, len(exp))
	for i, e := range exp {
		args[i] = Eval(env, e)
	}
	return fn(env, args...)
}

// Evaluate an expression
func Eval(env *types.Environment, exp types.Expression) types.Expression {
	// Only lists and symbols are "special",
	// the rest will just return
	switch exp := exp.(type) {
	case types.Symbol:
		return env.Lookup(exp)
	case types.List:
		car, cdr := types.Cons(exp)
		var (
			sym types.Symbol
			ok  bool
			fn  types.Function
		)
		if sym, ok = car.(types.Symbol); !ok {
			return ErrInvalidType
		}
		// Check special forms before environment bound symbols
		switch sym {
		case "quote":
			return applyQuote(env, cdr...)
		case "def":
			return applyDef(env, cdr...)
		case "if":
			return applyIf(env, cdr...)
		case "while":
			return applyWhile(env, cdr...)
		case "seq":
			return applySeq(env, cdr...)
		case "let":
			return applyLet(env, cdr...)
		case "fn", "lambda":
			return applyFn(env, cdr...)
		}
		if fn, ok = env.Lookup(sym).(types.Function); !ok {
			return ErrInvalidType
		}
		return apply(env, fn, cdr...)
	default:
		return exp
	}
}
