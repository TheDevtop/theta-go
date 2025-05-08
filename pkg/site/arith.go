package site

/*
	Theta list processor
	Specialized arithmetic
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

type number = interface{ int | float64 }

func arithCast[T number](expList ...types.Expression) []T {
	ret := make([]T, len(expList))
	for i, exp := range expList {
		switch exp := exp.(type) {
		case int:
			ret[i] = T(exp)
		case float64:
			ret[i] = T(exp)
		}
	}
	return ret
}

func arithAdd[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc += v
	}
	return acc
}

func arithMul[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc *= v
	}
	return acc
}

func arithSub[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc -= v
	}
	return acc
}

func arithDiv[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc /= v
	}
	return acc
}

func arithLesser[T number](set ...T) bool {
	for _, v := range set[1:] {
		if set[0] >= v {
			return false
		}
	}
	return true
}

func arithGreater[T number](set ...T) bool {
	for _, v := range set[1:] {
		if set[0] <= v {
			return false
		}
	}
	return true
}
