package types

/*
	Theta list processor
	Type utility functions
*/

// Extract car and cdr out of list
func Cons(list List) (Expression, List) {
	if len(list) < 1 {
		return nil, nil
	}
	if len(list) == 1 {
		return list[0], nil
	}
	return list[0], list[1:]
}

// Make a list from expressions
func MakeList(exp ...Expression) List {
	return List(exp)
}

// Reduce arguments back to list of symbols
func ReduceArgs(args []Symbol) List {
	list := make(List, len(args))
	for i, sym := range args {
		list[i] = sym
	}
	return list
}

// Check if all expressions are consistent with type
func IsConsistent[T Expression](expList ...Expression) bool {
	for _, exp := range expList {
		if _, ok := exp.(T); !ok {
			return false
		}
	}
	return true
}

// Cast all list expressions to type (best effort)
func Cast[T Expression](expList ...Expression) []T {
	var (
		nlist = make([]T, len(expList)) // Make sure that the numbers of items is the same
		ok    bool
		obj   T
	)

	for i, exp := range expList {
		if obj, ok = exp.(T); ok {
			nlist[i] = obj
		}
	}
	return nlist
}
