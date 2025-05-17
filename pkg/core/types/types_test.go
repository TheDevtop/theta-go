package types

import (
	"testing"
)

// Package axioms
var (
	deca Expression = int(10)
	list            = List{Keyword(":foo"), Keyword(":bar"), Keyword(":baz")}
	fn   Function   = func(_ *Environment, exp ...Expression) Expression {
		n := exp[0].(int)
		return n + 1
	}
)

func TestFn(t *testing.T) {
	// Axioms
	var expectA Expression = int(11)

	// Application
	output := fn(nil, deca)

	// Verification
	if output != expectA {
		t.Errorf("Output should be 11, is %v\n", output)
	}
}

func TestCons(t *testing.T) {
	// Axioms
	var expectA Expression = Keyword(":foo")

	// Application
	output, _ := Cons(list)

	// Verification
	if output != expectA {
		t.Errorf("Objects should be equal\n")
	}
}

func TestMakeList(t *testing.T) {
	// Axioms
	var (
		expectA Expression = Symbol("a")
		expectB Expression = Keyword(":ok")
	)

	// Application
	output := MakeList(expectA, expectB)

	// Verification
	if len(output) != 2 {
		t.Errorf("Output should be of length 2, is %d\n", len(output))
	}
	if output[0] != expectA || output[1] != expectB {
		t.Errorf("Objects should be equal\n")
	}
}

func TestIsConsistent(t *testing.T) {
	// Axioms
	var input = List{int(42), true, int(69), deca}

	// Application
	output := IsConsistent[int](input...)

	// Verification
	if output {
		t.Errorf("Output should be false\n")
	}
}

func TestCast(t *testing.T) {
	// Axioms
	var input = List{int(42), true, int(69), deca}

	// Application
	output := Cast[int](input...)

	// Verification
	if len(output) != len(input) {
		t.Errorf("Objects should be equal in length\n")
	}
	if output[1] != 0 {
		t.Errorf("Object should be 0, is %v\n", output[1])
	}
}
