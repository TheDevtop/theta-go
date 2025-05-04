package site

/*
	Theta virtual machine
	Site construction (initial environment)
*/

import "github.com/TheDevtop/theta-go/pkg/types"

var SiteTable = map[types.Symbol]types.Value{
	types.Symbol("=="):    siteEqual,
	types.Symbol("!="):    siteNequal,
	types.Symbol("len"):   siteLen,
	types.Symbol("list"):  siteList,
	types.Symbol("car"):   siteCar,
	types.Symbol("cdr"):   siteCdr,
	types.Symbol("apply"): siteApply,
	types.Symbol("mesg"):  siteMesg,
}
