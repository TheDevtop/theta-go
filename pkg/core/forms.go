package core

/*
	Theta list processor
	Special evaluation forms
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

// Pass evaluation on arguments
// (quote EXP...)
func applyQuote(_ *types.Environment, exp ...types.Expression) types.Expression {
	if len(exp) < 1 {
		return ErrInvalidArgs
	} else if len(exp) == 1 {
		return exp[0]
	} else {
		return exp
	}
}

// Define a new abstraction
// (def SYMBOL EXP)
func applyDef(env *types.Environment, exp ...types.Expression) types.Expression {
	var (
		sym types.Symbol
		ok  bool
	)
	if len(exp) != 2 {
		return ErrInvalidArgs
	}
	if sym, ok = exp[0].(types.Symbol); !ok {
		return ErrInvalidType
	}
	env.Modify(sym, Eval(env, exp[1]))
	return KeyOk
}

// Conditional evaluation
// (if CONDITIONAL THEN-EXP ELSE-EXP)
func applyIf(env *types.Environment, exp ...types.Expression) types.Expression {
	var (
		cond bool
		ok   bool
	)
	if len(exp) != 3 {
		return ErrInvalidArgs
	}
	if cond, ok = Eval(env, exp[0]).(bool); !ok {
		return ErrInvalidType
	}
	if cond {
		return Eval(env, exp[1])
	}
	return Eval(env, exp[2])
}

// While condition evaluate
// (while CONDITIONAL EXP)
func applyWhile(env *types.Environment, exp ...types.Expression) types.Expression {
	var (
		cond   bool
		ok     bool
		retExp types.Expression = nil
	)
	if len(exp) != 2 {
		return ErrInvalidArgs
	}
	for {
		if cond, ok = Eval(env, exp[0]).(bool); !ok {
			return ErrInvalidType
		}
		if !cond {
			return retExp
		} else {
			retExp = Eval(env, exp[1])
		}
	}
}

// Evaluate sequence
// (seq EXP...)
func applySeq(env *types.Environment, exp ...types.Expression) types.Expression {
	var retExp types.Expression
	for _, e := range exp {
		retExp = Eval(env, e)
	}
	return retExp
}

// Compose function
// (fn (SYMBOLS...) EXP)
func applyFn(_ *types.Environment, exp ...types.Expression) types.Expression {
	var (
		args []types.Symbol
		body types.Expression
		fn   types.Function
	)

	if len(exp) != 2 {
		return ErrInvalidArgs
	}
	body = exp[1]
	if list, ok := exp[0].(types.List); !ok {
		return ErrInvalidType
	} else {
		args = types.Cast[types.Symbol](list...)
	}

	fn = func(env *types.Environment, exp ...types.Expression) types.Expression {
		var (
			args = args
			body = body
			fenv = types.NewEnvironment(len(args))
		)
		if len(args) != len(exp) {
			return ErrInvalidArgs
		}
		for i, arg := range args {
			fenv.Modify(arg, exp[i])
		}
		fenv.Link(env)
		return Eval(fenv, body)
	}
	return fn
}
