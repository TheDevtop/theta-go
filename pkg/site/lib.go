package site

/*
	Theta virtual machine
	Site library
*/

import (
	"github.com/TheDevtop/theta-go/pkg/mce"
	"github.com/TheDevtop/theta-go/pkg/types"
)

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
)
