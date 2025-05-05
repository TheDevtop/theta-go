package sexp

/*
	Theta virtual machine
	S-expression parser
*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/TheDevtop/theta-go/pkg/types"
	"github.com/zyedidia/generic/stack"
)

const (
	keyTrue  types.Keyword = ":true"
	keyFalse types.Keyword = ":false"
	keyNil   types.Keyword = ":nil"
	keyFn    types.Keyword = ":fn"
)

func isKeyword(str string) (types.Keyword, bool) {
	if strings.HasPrefix(str, ":") && len(str) > 1 {
		return types.Keyword(str), true
	}
	return types.Keyword(str), false
}

func boolToKeyword(bit bool) types.Keyword {
	if bit {
		return keyTrue
	} else {
		return keyFalse
	}
}

func lex(str string) []string {
	rex := regexp.MustCompile(`"([^"\\]|\\.)*"|[()]|[^()\s]+`)
	return rex.FindAllString(str, -1)
}

func parse(token string) types.Value {
	if strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		return token
	}
	if i, err := strconv.Atoi(token); err == nil {
		return int32(i)
	}
	if f, err := strconv.ParseFloat(token, 32); err == nil {
		return float32(f)
	}
	if k, ok := isKeyword(token); ok {
		if k == keyTrue {
			return true
		}
		if k == keyFalse {
			return false
		}
		if k == keyNil {
			return nil
		}
		return k
	}
	return types.Symbol(token)
}

// Encode s-expression, returns string
func Marshal(val types.Value) string {
	ret := ""
	switch val := val.(type) {
	case nil:
		ret = string(keyNil)
	case bool:
		ret = string(boolToKeyword(val))
	case string:
		ret = val
	case int32:
		ret = fmt.Sprintf("%d", val)
	case float32:
		ret = fmt.Sprintf("%f", val)
	case types.Symbol:
		ret = string(val)
	case types.Keyword:
		ret = string(val)
	case types.Lambda:
		ret = string(keyFn)
	case types.List:
		bucket := make([]string, 0, 8)
		for _, v := range val {
			bucket = append(bucket, Marshal(v))
		}
		ret = "(" + strings.Join(bucket, " ") + ")"
	}
	return ret
}

// Decode s-expression from string
func Unmarshal(str string) types.Value {
	var (
		tokens = lex(str)
		ret    = make(types.List, 0)
		sptr   = stack.New[types.List]()
	)

	for _, tok := range tokens {
		switch tok {
		case "(":
			sptr.Push(ret)
			ret = make(types.List, 0)
		case ")":
			ret = append(sptr.Pop(), ret)
		default:
			ret = append(ret, parse(tok))
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret[0]
}
