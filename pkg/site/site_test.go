package site

import (
	"testing"

	"github.com/TheDevtop/theta-go/pkg/core"
	"github.com/TheDevtop/theta-go/pkg/core/sexp"
)

// Repl for site testing
func repl(str string) string {
	return sexp.Encode(core.Eval(DefaultSite, sexp.Decode(str)))
}

func TestAdd(t *testing.T) {
	// Axioms
	var (
		input  = "(+ 8 2)"
		expect = "10"
	)

	// Application
	output := repl(input)

	// Verification
	if output != expect {
		t.Errorf("Objects should be equal\n")
	}
}

func TestEqual(t *testing.T) {
	// Axioms
	var (
		input  = "(= 10 deca)"
		expect = ":false"
	)

	// Application
	output := repl(input)

	// Verification
	if output != expect {
		t.Errorf("Objects should be equal\n")
	}
}

func TestCar(t *testing.T) {
	// Axioms
	var (
		input  = "(car (quote (a b c)))"
		expect = "a"
	)

	// Application
	output := repl(input)

	// Verification
	if output != expect {
		t.Errorf("Objects should be equal\n")
	}
}

func TestCdr(t *testing.T) {
	// Axioms
	var (
		input  = "(cdr (quote (a b c)))"
		expect = "(b c)"
	)

	// Application
	output := repl(input)

	// Verification
	if output != expect {
		t.Errorf("Objects should be equal\n")
	}
}
