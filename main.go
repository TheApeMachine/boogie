package main

import (
	"bufio"
	"fmt"
	"github.com/theapemachine/boogie/boogie"
	"os"
)

func main() {
	lic := make(chan string)
	lexer := boogie.NewLexer(lic)
	reader := bufio.NewReader(os.Stdin)

	go lexer.Run()

	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		lexer.In <- text
	}
}
