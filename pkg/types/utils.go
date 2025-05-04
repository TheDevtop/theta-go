package types

/*
	Theta virtual machine
	Utility functions for types
*/

import (
	"fmt"
)

const (
	KeyOk  Keyword = ":ok"
	KeyErr Keyword = ":err"
)

// Extract car and cdr out of list
func Cons(list List) (Value, List) {
	if len(list) < 1 {
		return nil, nil
	}
	if len(list) == 1 {
		return list[0], nil
	}
	return list[0], list[1:]
}

// Merge car and cdr into list
func Merge(car Value, cdr List) List {
	return append(List{car}, cdr...)
}

// Create message with status keyword
func Message(stat Keyword, mesg string) Value {
	return Value(List{stat, fmt.Sprintf("\"%s\"", mesg)})
}

// Check if all list objects are consistent with type
func IsConsistent[T any](list List) bool {
	for _, exp := range list {
		if _, ok := exp.(T); !ok {
			return false
		}
	}
	return true
}
