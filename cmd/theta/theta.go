package main

/*
	Theta list processor
	Server-side program
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/TheDevtop/theta-go/pkg/core"
	"github.com/TheDevtop/theta-go/pkg/core/sexp"
	"github.com/TheDevtop/theta-go/pkg/core/types"
	"github.com/TheDevtop/theta-go/pkg/csio"
	"github.com/TheDevtop/theta-go/pkg/site"
)

func handleStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, sexp.Encode(core.KeyOk))
}

func handleEval(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		buf []byte
		exp types.Expression
	)

	if buf, err = io.ReadAll(r.Body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exp = core.Eval(site.DefaultSite, sexp.Decode(string(buf)))
	fmt.Fprint(w, sexp.Encode(exp))
}

func main() {

	http.HandleFunc(csio.PathEval, handleEval)
	http.HandleFunc(csio.PathStatus, handleStatus)

	log.Println("Theta list processor")
	if err := http.ListenAndServe(csio.DefaultPort, nil); err != nil {
		log.Println(err)
	}
}
