package sexp

/*
	Theta list processor
	S-Expression parser
*/

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/TheDevtop/theta-go/pkg/core/types"
)

const (
	keyBin   types.Keyword = ":bin"
	symTrue  types.Symbol  = "true"
	symFalse types.Symbol  = "false"
)

// Check if string is keyword
func isKeyword(str string) (types.Keyword, bool) {
	if strings.HasPrefix(str, ":") && len(str) > 1 {
		return types.Keyword(str), true
	}
	return types.Keyword(str), false
}

// Convert boolean to symbol representation
func boolToKeyword(bit bool) types.Symbol {
	if bit {
		return symTrue
	} else {
		return symFalse
	}
}

// Tokenize input string via special regex
func lex(str string) []string {
	rex := regexp.MustCompile(`"([^"\\]|\\.)*"|[()]|[^()\s]+`)
	return rex.FindAllString(str, -1)
}

// Decode anything that is not a list
func parse(token string) types.Expression {
	if strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		return token
	}
	if i, err := strconv.Atoi(token); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(token, 32); err == nil {
		return f
	}
	if k, ok := isKeyword(token); ok {
		return k
	}
	switch token {
	case "nil":
		return nil
	case "true":
		return true
	case "false":
		return false
	default:
		return types.Symbol(token)
	}
}

// Decode s-expression from string
func Decode(str string) types.Expression {
	var (
		tokens = lex(str)
		ret    = types.List{}
		stack  = make([]types.List, 0)
	)

	for _, tok := range tokens {
		switch tok {
		case "(":
			stack = append(stack, ret)
			ret = types.List{}
		case ")":
			if len(stack) == 0 {
				ret = append(types.List{}, ret)
			} else {
				ret = append(stack[len(stack)-1], ret)
				stack = stack[:len(stack)-1]
			}
		default:
			ret = append(ret, parse(tok))
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret[0]
}

// Encode s-expression to string
func Encode(exp types.Expression) string {
	ret := ""
	switch exp := exp.(type) {
	case nil:
		ret = string("nil")
	case bool:
		ret = string(boolToKeyword(exp))
	case string:
		ret = exp
	case int:
		ret = strconv.Itoa(exp)
	case float64:
		ret = strconv.FormatFloat(exp, 'f', 4, 32)
	case types.Symbol:
		ret = string(exp)
	case types.Keyword:
		ret = string(exp)
	case types.Function:
		ret = string(keyBin)
	case types.Procedure:
		ret = string(keyBin)
	case types.List:
		bucket := make([]string, 0, 8)
		for _, e := range exp {
			bucket = append(bucket, Encode(e))
		}
		ret = "(" + strings.Join(bucket, " ") + ")"
	}
	return ret
}
