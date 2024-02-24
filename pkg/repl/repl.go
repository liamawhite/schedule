package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/token"
)

const PROMPT = ">> "
const WELCOME = "Welcome to the Schedule REPL! Type a schedule to parse it."

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    fmt.Println(WELCOME)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }

}
