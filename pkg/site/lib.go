package site

/*
	Theta virtual machine
	Site library
*/

import (
	"github.com/TheDevtop/theta-go/pkg/mce"
	"github.com/TheDevtop/theta-go/pkg/types"
)

var errArith = types.Message(types.KeyErr, "Arithmetic error")

// Library functions as typed lambdas
var (
	siteEqual types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) != 2 {
			return mce.ErrInvalidArgs
		}
		return args[0] == args[1]
	}
	siteNequal types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) != 2 {
			return mce.ErrInvalidArgs
		}
		return args[0] != args[1]
	}
	siteAdd types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) < 1 {
			return mce.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return genAdd(genCast[int32](args))
		case float32:
			return genAdd(genCast[float32](args))
		}
		return errArith
	}
	siteMul types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) < 1 {
			return mce.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return genMul(genCast[int32](args))
		case float32:
			return genMul(genCast[float32](args))
		}
		return errArith
	}
	siteSub types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) < 1 {
			return mce.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return genSub(genCast[int32](args))
		case float32:
			return genSub(genCast[float32](args))
		}
		return errArith
	}
	siteDiv types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) < 1 {
			return mce.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int32:
			return genDiv(genCast[int32](args))
		case float32:
			return genDiv(genCast[float32](args))
		}
		return errArith
	}
	siteAnd types.Lambda = func(args types.List, env *types.Environment) types.Value {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return mce.ErrInvalidArgs
			} else if !res {
				return false
			}
		}
		return true
	}
	siteOr types.Lambda = func(args types.List, env *types.Environment) types.Value {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return mce.ErrInvalidArgs
			} else if res {
				return true
			}
		}
		return false
	}
	siteXor types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) != 2 {
			return mce.ErrInvalidArgs
		}
		nargs := genCast[bool](args)
		return nargs[0] != nargs[1]
	}
	siteLen types.Lambda = func(args types.List, env *types.Environment) types.Value {
		return int32(len(args))
	}
	siteList types.Lambda = func(args types.List, env *types.Environment) types.Value {
		return args
	}
	siteCar types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) != 1 {
			return mce.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); ok {
			car, _ := types.Cons(arg)
			return car
		}
		return mce.ErrInvalidArgs
	}
	siteCdr types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) != 1 {
			return mce.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); ok {
			_, cdr := types.Cons(arg)
			return cdr
		}
		return mce.ErrInvalidArgs
	}
	siteApply types.Lambda = func(args types.List, env *types.Environment) types.Value {
		var (
			fn types.Lambda
			ok bool
		)
		if len(args) != 2 {
			return mce.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Lambda); !ok {
			return types.Message(types.KeyErr, "Invalid application type")
		}
		return fn(types.List{args[1]}, env)
	}
	siteMesg types.Lambda = func(args types.List, env *types.Environment) types.Value {
		if len(args) == 1 {
			if msg, ok := args[0].(string); ok {
				return types.List{types.KeyOk, msg}
			}
			return mce.ErrInvalidArgs
		} else if len(args) == 2 {
			key, okx := args[0].(types.Keyword)
			msg, oky := args[1].(string)
			if okx && oky {
				return types.List{key, msg}
			}
			return mce.ErrInvalidArgs
		}
		return mce.ErrInvalidArgs
	}
)
