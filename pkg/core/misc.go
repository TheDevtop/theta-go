package core

/*
	Theta list processor
	Misc
*/

import "github.com/TheDevtop/theta-go/pkg/core/types"

var (
	// Operational keywords
	KeyOk  = types.Keyword(":ok")
	KeyErr = types.Keyword(":err")

	// Error messages
	ErrInvalidArgs = Message(KeyErr, "Invalid arguments")
	ErrInvalidType = Message(KeyErr, "Invalid type or cast")
)

// Create message with keyword as label
func Message(key types.Keyword, mesg string) types.Expression {
	return types.List{key, "\"" + mesg + "\""}
}

func Call(env *types.Environment, fn types.Function, exp ...types.Expression) types.Expression {
	var fnenv = types.NewEnvironment(len(fn.Args))
	if len(exp) != len(fn.Args) {
		return ErrInvalidArgs
	}
	for i, arg := range fn.Args {
		fnenv.Modify(arg, exp[i])
	}
	fnenv.Link(env)
	return Eval(fnenv, fn.Body)
}
