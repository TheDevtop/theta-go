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

// Construct a lambda function
func Lambda(args []types.Symbol, body types.Expression) types.Procedure {
	var proc types.Procedure = func(env *types.Environment, exp ...types.Expression) types.Expression {
		var (
			args  = args
			body  = body
			fnenv = types.NewEnvironment(len(args))
		)
		if len(exp) != len(args) {
			return ErrInvalidArgs
		}
		for i, arg := range args {
			fnenv.Modify(arg, exp[i])
		}
		fnenv.Link(env)
		return Eval(fnenv, body)
	}
	return proc
}
