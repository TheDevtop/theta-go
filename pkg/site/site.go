package site

/*
	Theta list processor
	Site construction
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

var DefaultSite *types.Environment = &types.Environment{
	Parent: nil,
	Table: map[types.Symbol]types.Expression{
		types.Symbol("="):      siteEqual,
		types.Symbol("!="):     siteNequal,
		types.Symbol("+"):      siteAdd,
		types.Symbol("*"):      siteMul,
		types.Symbol("-"):      siteSub,
		types.Symbol("/"):      siteDiv,
		types.Symbol("<"):      siteLesser,
		types.Symbol(">"):      siteGreater,
		types.Symbol("and"):    siteAnd,
		types.Symbol("or"):     siteOr,
		types.Symbol("xor"):    siteXor,
		types.Symbol("not"):    siteNot,
		types.Symbol("len"):    siteLen,
		types.Symbol("list"):   siteList,
		types.Symbol("car"):    siteCar,
		types.Symbol("cdr"):    siteCdr,
		types.Symbol("rev"):    siteRev,
		types.Symbol("mem"):    siteMem,
		types.Symbol("append"): siteAppend,
		types.Symbol("apply"):  siteApply,
		types.Symbol("mesg"):   siteMesg,
		types.Symbol("map"):    siteMap,
		types.Symbol("filter"): siteFilter,
		types.Symbol("concat"): siteConcat,
		types.Symbol("printf"): sitePrintf,
		types.Symbol("nil?"):   siteIsNil,
		types.Symbol("bool?"):  siteIsBoolean,
		types.Symbol("str?"):   siteIsString,
		types.Symbol("int?"):   siteIsInteger,
		types.Symbol("float?"): siteIsFloating,
		types.Symbol("sym?"):   siteIsSymbol,
		types.Symbol("key?"):   siteIsKeyword,
		types.Symbol("fn?"):    siteIsFunction,
		types.Symbol("proc?"):  siteIsProcedure,
		types.Symbol("atom?"):  siteIsAtom,
		types.Symbol("list?"):  siteIsList,
		types.Symbol("unfn"):   siteUnfunction,
	},
}
