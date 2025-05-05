package main

/*
	Theta virtual machine
	Client program and repl
*/

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TheDevtop/theta-go/pkg/csio"
	"github.com/chzyer/readline"
)

// Post an expression to the Theta server
func postExp(addr string, exp string) (string, error) {
	var (
		body = strings.NewReader(exp)
		err  error
		res  *http.Response
		buf  []byte
	)
	addr = "http://" + addr + csio.PathEval
	if res, err = http.Post(addr, "text/x-elisp", body); err != nil {
		return "", err
	}
	if buf, err = io.ReadAll(res.Body); err != nil {
		return "", err
	}
	return string(buf), nil
}

// Client-side repl
func repl(addr string) {
	var (
		rl  *readline.Instance
		ln  string
		err error
	)

	if rl, err = readline.New("Θ> "); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		if ln, err = rl.Readline(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if ln == "(exit)" {
			os.Exit(0)
		}
		if ln, err = postExp(addr, ln); err != nil {
			fmt.Println(err)
		}
		fmt.Println(ln)
	}
}

func main() {
	var (
		flagAddr = flag.String("a", "127.0.0.1"+csio.DefaultPort, "Specify address")
		flagExp  = flag.String("e", "", "Specify expression")
	)

	flag.Parse()

	// If expression is specified post a single expression and exit,
	// if not start the repl.
	if *flagExp != "" {
		if res, err := postExp(*flagAddr, *flagExp); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(res)
			os.Exit(0)
		}
	} else {
		repl(*flagAddr)
	}
}
