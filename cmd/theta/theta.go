package main

/*
	Theta virtual machine
	Server program
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/TheDevtop/theta-go/pkg/csio"
	"github.com/TheDevtop/theta-go/pkg/mce"
	"github.com/TheDevtop/theta-go/pkg/sexp"
	"github.com/TheDevtop/theta-go/pkg/site"
	"github.com/TheDevtop/theta-go/pkg/types"
)

var env *types.Environment

func handleEval(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		buf []byte
		val types.Value
	)

	if buf, err = io.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	val = mce.Eval(sexp.Unmarshal(string(buf)), env)
	fmt.Fprint(w, sexp.Marshal(val))
}

func main() {
	env = types.InitEnvironment(site.SiteTable)
	http.HandleFunc(csio.PathEval, handleEval)

	log.Println("(Θ list processor)")
	if err := http.ListenAndServe(csio.DefaultPort, nil); err != nil {
		log.Println(err)
	}
}
