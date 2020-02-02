package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"moose/lexer"
	"moose/token"
)

const PROMPT = ">_"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		l := lexer.New("console", strings.NewReader(scanner.Text()))
		for t := l.Next(); t.Type != token.EOF; t = l.Next() {
			fmt.Printf("%+v\n", t)
		}
	}
}
