package site

/*
	Theta list processor
	Site implementation
*/

import (
	"fmt"
	"slices"
	"strings"

	"github.com/TheDevtop/theta-go/pkg/core"
	"github.com/TheDevtop/theta-go/pkg/core/types"
)

var (
	siteEqual types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		return args[0] == args[1]
	}
	siteNequal types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		return args[0] != args[1]
	}
	siteAdd types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithAdd(arithCast[int](args...)...)
		case float64:
			return arithAdd(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteMul types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithMul(arithCast[int](args...)...)
		case float64:
			return arithMul(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteSub types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithSub(arithCast[int](args...)...)
		case float64:
			return arithSub(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteDiv types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 1 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithDiv(arithCast[int](args...)...)
		case float64:
			return arithDiv(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteLesser types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithLesser(arithCast[int](args...)...)
		case float64:
			return arithLesser(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteGreater types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		switch args[0].(type) {
		case int:
			return arithGreater(arithCast[int](args...)...)
		case float64:
			return arithGreater(arithCast[float64](args...)...)
		default:
			return core.ErrInvalidType
		}
	}
	siteAnd types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return core.ErrInvalidType
			} else if !res {
				return false
			}
		}
		return true
	}
	siteOr types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		for _, arg := range args {
			if res, ok := arg.(bool); !ok {
				return core.ErrInvalidType
			} else if res {
				return true
			}
		}
		return false
	}
	siteXor types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		nargs := types.Cast[bool](args...)
		return nargs[0] != nargs[1]
	}
	siteNot types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		nargs := types.Cast[bool](args...)
		return !nargs[0]
	}
	siteLen types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		return len(args)
	}
	siteList types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return args
	}
	siteCar types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
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
	siteCdr types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
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
	siteRev types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if arg, ok := args[0].(types.List); !ok {
			return core.ErrInvalidType
		} else {
			slices.Reverse(arg)
			return arg
		}
	}
	siteMem types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			list types.List
			ok   bool
		)
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		if list, ok = args[1].(types.List); !ok {
			return core.ErrInvalidType
		}
		return slices.Contains(list, args[0])
	}
	siteAppend types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		if list, ok := args[0].(types.List); !ok {
			return core.ErrInvalidType
		} else {
			rest := args[1:]
			return append(list, rest...)
		}
	}
	siteApply types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn types.Procedure
			ok bool
		)
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Procedure); !ok {
			return core.ErrInvalidType
		}
		return fn(env, args[1:]...)
	}
	siteMesg types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
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
	siteMap types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn   types.Procedure
			list types.List
			ok   bool
		)
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Procedure); !ok {
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
	siteFilter types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fn      types.Procedure
			inList  types.List
			outList types.List
			ok      bool
		)
		if len(args) != 2 {
			return core.ErrInvalidArgs
		}
		if fn, ok = args[0].(types.Procedure); !ok {
			return core.ErrInvalidType
		}
		if inList, ok = args[1].(types.List); !ok {
			return core.ErrInvalidType
		}
		outList = make(types.List, 0, len(inList))

		for _, exp := range inList {
			bit, ok := fn(env, exp).(bool)
			if ok && bit {
				outList = append(outList, exp)
			}
		}
		return outList
	}
	siteConcat types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var nargs []string = make([]string, len(args))
		if !types.IsConsistent[string](args...) {
			return core.ErrInvalidType
		}
		for i, e := range args {
			nargs[i] = unmapQuotes(e.(string))
		}
		return mapQuotes(strings.Join(nargs, " "))
	}
	sitePrintf types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		var (
			fmtStr string
			nargs  []any
		)
		if len(args) < 2 {
			return core.ErrInvalidArgs
		}
		if !types.IsConsistent[string](args[0]) {
			return core.ErrInvalidType
		}
		fmtStr = unmapQuotes(args[0].(string))
		nargs = make([]any, len(args)-1)
		for i, e := range args[1:] {
			if ce, ok := e.(string); ok {
				nargs[i] = unmapQuotes(ce)
			} else {
				nargs[i] = e
			}
		}
		return mapQuotes(fmt.Sprintf(fmtStr, nargs...))
	}
	siteIsNil types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		for _, arg := range args {
			if arg != nil {
				return false
			}
		}
		return true
	}
	siteIsBoolean types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[bool](args...)
	}
	siteIsString types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[string](args...)
	}
	siteIsInteger types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[int](args...)
	}
	siteIsFloating types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[float64](args...)
	}
	siteIsSymbol types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[types.Symbol](args...)
	}
	siteIsKeyword types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[types.Keyword](args...)
	}
	siteIsFunction types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[types.Function](args...)
	}
	siteIsProcedure types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[types.Procedure](args...)
	}
	siteIsAtom types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if !types.IsConsistent[types.List](args...) && !types.IsConsistent[types.Procedure](args...) {
			return true
		}
		return false
	}
	siteIsList types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		return types.IsConsistent[types.List](args...)
	}
	siteUnfunction types.Procedure = func(env *types.Environment, args ...types.Expression) types.Expression {
		if len(args) != 1 {
			return core.ErrInvalidArgs
		}
		if !types.IsConsistent[types.Function](args[0]) {
			return core.ErrInvalidType
		}
		fn := args[0].(types.Function)

		return types.List{types.Symbol("fn"), types.ReduceArgs(fn.Args), fn.Body}
	}
)
