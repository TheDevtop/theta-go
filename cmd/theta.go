package main

import (
	"fmt"
	"os"

	"github.com/TheDevtop/theta-go/pkg/mce"
	"github.com/TheDevtop/theta-go/pkg/sexp"
)

func main() {
	if len(os.Args) != 2 {
		panic("Theta needs an argument to evaluate")
	}
	exp := mce.Eval(sexp.Unmarshal(os.Args[1]), mce.New())
	fmt.Println(sexp.Marshal(exp))
}
