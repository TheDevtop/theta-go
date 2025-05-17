package types

import (
	"testing"
)

func TestLookup(t *testing.T) {
	// Axioms
	var env = NewEnvironment(1)
	env.Modify(Symbol("deca"), int(10))

	// Application
	outputA := env.Lookup(Symbol("error"))
	outputB := env.Lookup(Symbol("deca"))

	// Verification
	if outputA != nil || outputB != 10 {
		t.Errorf("Test failed\n")
	}
}

func TestDelete(t *testing.T) {
	// Axioms
	var env = NewEnvironment(1)
	env.Modify(Symbol("deca"), int(10))

	// Application
	env.Modify(Symbol("deca"), nil)

	// Verification
	if env.Table[Symbol("deca")] != nil {
		t.Errorf("Object should be nil \n")
	}
}
