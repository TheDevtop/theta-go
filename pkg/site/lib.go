package site

/*
	Theta virtual machine
	Site library
*/

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
	// siteAdd types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {}
	// siteMul types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {}
	// siteSub types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {}
	// siteDiv types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {}
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
	siteLen types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		return int32(len(args))
	}
	siteList types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		return args
	}
	siteCar types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); ok {
			car, _ := types.Cons(arg)
			return car
		}
		return core.ErrInvalidType
	}
	siteCdr types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); ok {
			_, cdr := types.Cons(arg)
			return cdr
		}
		return core.ErrInvalidType
	}
	siteApply types.Function = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn types.Function
			ok bool
		)
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Function); !ok {
			return core.ErrInvalidType
		}
		return fn(env, types.List{args[1]})
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

		// If not a list: map single value, return single.
		// If a list: map values, return list of mapped values.
		if list, ok = args[1].(types.List); !ok {
			return fn(env, types.List{args[1]})
		}
		for i, exp := range list {
			list[i] = fn(env, types.List{exp})
		}
		return list
	}
)

// Library functions as typed lambdas
// var (
// 	siteEqual types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	siteNequal types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteAdd types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		if len(args) < 1 {
// 			return mce.ErrInvalidArgs
// 		}
// 		switch args[0].(type) {
// 		case int32:
// 			return genAdd(genCast[int32](args))
// 		case float32:
// 			return genAdd(genCast[float32](args))
// 		}
// 		return errArith
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
// 	siteAnd types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteOr types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteXor types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	siteLen types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		return int32(len(args))
// 	}
// 	siteList types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 		return args
// 	}
// 	siteCar types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteCdr types.Lambda = func(args types.List, env *types.Environment) types.Value {

// 	}
// 	siteA
// 	}
// 	siteMap types.Lambda = func(args types.List, env *types.Environment) types.Value {
// 	}
// )
