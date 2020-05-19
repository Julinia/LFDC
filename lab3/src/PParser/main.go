package main

import (
	"PParser/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Type any commands that are specified in readme\n")
	repl.Start(os.Stdin, os.Stdout)
}
