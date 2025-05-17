package core

import (
	"testing"

	"github.com/TheDevtop/theta-go/pkg/core/types"
)

func TestLambda(t *testing.T) {
	// Axioms
	expect := KeyOk
	fn := Lambda(
		[]types.Symbol{types.Symbol("i")},
		types.List{
			types.Symbol("if"),
			types.Symbol("i"),
			KeyOk,
			KeyErr,
		},
	)

	// Application
	output := fn(nil, true)

	// Verification
	if output != expect {
		t.Fatalf("Output should be true\n")
	}
}

func TestEval(t *testing.T) {
	// Axioms

	// Application

	// Verification

}
