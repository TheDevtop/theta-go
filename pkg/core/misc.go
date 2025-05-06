package core

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
