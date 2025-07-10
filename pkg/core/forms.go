package core

/*
	Theta list processor
	Special evaluation forms
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

// Pass evaluation on arguments
// (quote EXP)
func applyQuote(_ *types.Environment, exp ...types.Expression) types.Expression {
	if len(exp) != 1 {
		return ErrInvalidArgs
	}
	return exp[0]
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
	return sym
}

// Conditional evaluation
// (if COND THEN-EXP ELSE-EXP)
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
// (while COND EXP...)
func applyWhile(env *types.Environment, exp ...types.Expression) types.Expression {
	var (
		ok     bool
		cond   bool
		retExp types.Expression
	)
	if len(exp) != 2 {
		return ErrInvalidArgs
	}
	for {
		if cond, ok = Eval(env, exp[0]).(bool); !ok {
			return ErrInvalidType
		} else if !cond {
			return retExp
		}
		retExp = applySeq(env, exp[1:]...)
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

// Let expression
// (let ((SYM INIT-EXP)...) EXP...)
func applyLet(env *types.Environment, exp ...types.Expression) types.Expression {
	if len(exp) < 2 {
		return ErrInvalidArgs
	}
	var (
		bodyExp    = exp[1:]
		sePairList types.List
		ok         bool
		nenv       *types.Environment
	)
	if sePairList, ok = exp[0].(types.List); !ok {
		return ErrInvalidType
	}
	nenv = types.NewEnvironment(len(sePairList))
	for _, sePair := range sePairList {
		// Extract symbol/expression pair out of the list
		// Extract symbol out of pair
		if sePair, ok := sePair.(types.List); !ok {
			return ErrInvalidType
		} else if len(sePair) != 2 {
			return ErrInvalidArgs
		} else if sym, ok := sePair[0].(types.Symbol); !ok {
			return ErrInvalidType
		} else {
			// Bind symbol with the evaluation of the expression
			nenv.Modify(sym, Eval(env, sePair[1]))
		}
	}
	nenv.Link(env)
	return applySeq(nenv, bodyExp...)
}

// Compose function via Lambda()
// (fn (SYMBOLS...) EXP)
func applyFn(_ *types.Environment, exp ...types.Expression) types.Expression {
	var fn types.Function

	if len(exp) != 2 {
		return ErrInvalidArgs
	}
	fn.Body = exp[1]
	if list, ok := exp[0].(types.List); !ok {
		return ErrInvalidType
	} else {
		fn.Args = types.Cast[types.Symbol](list...)
	}
	return fn
}
