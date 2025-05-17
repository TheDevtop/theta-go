package sexp

import (
	"testing"

	"github.com/TheDevtop/theta-go/pkg/core/types"
)

func TestDecode(t *testing.T) {
	// Axioms
	var (
		input   = "(a b (c))"
		expectA = types.List{types.Symbol("a"), types.Symbol("b"), types.List{types.Symbol("c")}}
	)

	// Application
	output := Decode(input)

	// Verification
	list, ok := output.(types.List)
	if !ok {
		t.Errorf("Output should be castable to list\n")
	}
	if len(list) != len(expectA) {
		t.Errorf("Output and expectation should be of equal lenght\n")
	}
	if list[0] != expectA[0] {
		t.Errorf("Objects should be equal\n")
	}
	if obj, ok := list[2].(types.List); !ok {
		t.Errorf("Object shoud be of type list, got: %v\n", obj)
	}
}

func TestEncode(t *testing.T) {
	// Axioms
	var (
		input          = types.List{types.Symbol("a"), types.Symbol("b"), types.List{types.Symbol("c")}}
		expectA string = "(a b c)"
		expectB string = "(a b (c))"
	)

	// Application
	output := Encode(input)

	// Verification
	if output == expectA || output != expectB {
		t.Errorf("Output should be: %s\n", expectB)
	}
}
