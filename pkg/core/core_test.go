package core

import (
	"testing"

	"github.com/TheDevtop/theta-go/pkg/core/types"
)

func TestLambda(t *testing.T) {
	// Axioms
	expect := KeyOk

	fn := types.Function{
		Args: []types.Symbol{types.Symbol("i")},
		Body: types.List{
			types.Symbol("if"),
			types.Symbol("i"),
			KeyOk,
			KeyErr,
		},
	}

	// Application
	output := Call(nil, fn, true)

	// Verification
	if output != expect {
		t.Fatalf("Output should be true, result: %v\n", output)
	}
}

func TestEval(t *testing.T) {
	// Axioms

	// Application

	// Verification

}
