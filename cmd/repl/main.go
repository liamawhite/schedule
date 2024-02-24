package main

import (
	"os"

	"github.com/liamawhite/schedule/pkg/repl"
)

func main() {
    repl.Start(os.Stdin, os.Stdout)
}

