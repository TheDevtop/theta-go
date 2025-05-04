package mce

import "github.com/TheDevtop/theta-go/pkg/sexp"

type Lambda func(sexp.List, *Environment) sexp.Value

func (Lambda) String() string {
	return ":fn"
}
