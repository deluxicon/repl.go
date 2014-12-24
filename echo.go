package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()
	fmt.Println(scanner.Text())

	repl()
}

func main() {
	repl()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
