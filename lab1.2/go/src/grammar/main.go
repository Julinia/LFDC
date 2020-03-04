package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"./lexer"
)

const PROMPT = "$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func main() {

	fmt.Printf("Type a command\n")
	Start(os.Stdin, os.Stdout)
}
