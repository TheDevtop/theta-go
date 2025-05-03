package main

import (
	"fmt"

	"github.com/TheDevtop/theta-go/pkg/sexp"
)

func main() {
	val := sexp.Unmarshal(`(write :stdout (concat "Hello, " "world!"))`)
	fmt.Printf("%v\n", val)
	fmt.Println(sexp.Marshal(val))
}
