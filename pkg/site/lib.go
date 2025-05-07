package site

import (
	"github.com/TheDevtop/theta-go/pkg/core"
	"github.com/TheDevtop/theta-go/pkg/core/types"
)

var (
	siteEqual types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		return args[0] == args[1]
	}
	siteNequal types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		return args[0] != args[1]
	}
	siteAdd types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithAdd(types.Cast[int32](args...)...)
		case float32:
			return arithAdd(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteMul types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithMul(types.Cast[int32](args...)...)
		case float32:
			return arithMul(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteSub types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithSub(types.Cast[int32](args...)...)
		case float32:
			return arithSub(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteDiv types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithDiv(types.Cast[int32](args...)...)
		case float32:
			return arithDiv(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteLesser types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithLesser(types.Cast[int32](args...)...)
		case float32:
			return arithLesser(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteGreater types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return arithGreater(types.Cast[int32](args...)...)
		case float32:
			return arithGreater(types.Cast[float32](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteAnd types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return core.ErrInvalidType
			} else if !res {
				return false
			}
		}
		return true
	}
	siteOr types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return core.ErrInvalidType
			} else if res {
				return true
			}
		}
		return false
	}
	siteXor types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		nargs := types.Cast[bool](args...)
		return nargs[0] != nargs[1]
	}
	siteNot types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		nargs := types.Cast[bool](args...)
		return !nargs[0]
	}
	siteLen types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		return int32(len(args))
	}
	siteList types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		return args
	}
	siteCar types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); !ok {
			return core.ErrInvalidType
		} else {
			car, _ := types.Cons(arg)
			return car
		}
	}
	siteCdr types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); !ok {
			return core.ErrInvalidType
		} else {
			_, cdr := types.Cons(arg)
			return cdr
		}
	}
	siteApply types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn types.Function
			ok bool
		)
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Function); !ok {
			return core.ErrInvalidType
		}
		return fn(env, args[1:]...)
	}
	siteMesg types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) == 1 {
			if msg, ok := args[0].(string); ok {
				return types.List{core.KeyOk, msg}
			}
			return core.ErrInvalidType
		} else if len(args) == 2 {
			key, okx := args[0].(types.Keyword)
			msg, oky := args[1].(string)
			if okx && oky {
				return types.List{key, msg}
			}
			return core.ErrInvalidType
		}
		return core.ErrInvalidArgs
	}
	siteMap types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn   types.Function
			list types.List
			ok   bool
		)
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Function); !ok {
			return core.ErrInvalidType
		}
		if list, ok = args[1].(types.List); !ok {
			return core.ErrInvalidType
		}

		for i, exp := range list {
			list[i] = fn(env, exp)
		}
		return list
	}
)

// Library functions as typed lambdas
// siteAdd types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteMul types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		if len(args) < 1 {
// 			return mce.ErrInvalidArgs
// 		}
// 		switch args[0].(type) {
// 		case int32:
// 			return genMul(genCast[int32](args))
// 		case float32:
// 			return genMul(genCast[float32](args))
// 		}
// 		return errArith
// 	}
// 	siteSub types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		if len(args) < 1 {
// 			return mce.ErrInvalidArgs
// 		}
// 		switch args[0].(type) {
// 		case int32:
// 			return genSub(genCast[int32](args))
// 		case float32:
// 			return genSub(genCast[float32](args))
// 		}
// 		return errArith
// 	}
// 	siteDiv types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		if len(args) < 1 {
// 			return mce.ErrInvalidArgs
// 		}
// 		switch args[0].(type) {
// 		case int32:
// 			return genDiv(genCast[int32](args))
// 		case float32:
// 			return genDiv(genCast[float32](args))
// 		}
// 		return errArith
// 	}
