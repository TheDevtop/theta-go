package sexp

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
		return k
	}
	return types.Symbol(token)
}

// Encode s-expression, returns string
func Marshal(val types.Value) string {
	ret := ""
	switch val.(type) {
	case nil:
		ret = string(types.KeyNil)
	case bool:
		_, v := val.(bool)
		ret = string(boolToKeyword(v))
	case string:
		ret = fmt.Sprint(val)
	case int32:
		ret = fmt.Sprintf("%d", val)
	case float32:
		ret = fmt.Sprintf("%f", val)
	case types.Symbol, types.Keyword:
		ret = fmt.Sprintf("%s", val)
	case types.Lambda:
		ret = string(keyFn)
	case types.List:
		bucket := make([]string, 0, 8)
		for _, v := range val.(types.List) {
			bucket = append(bucket, Marshal(v))
		}
		ret = fmt.Sprintf("(%s)", strings.Join(bucket, " "))
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
		return types.Value(nil)
	}
	return ret[0]
}
