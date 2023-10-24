package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "github.com/chettriyuvraj/lox-crafting-interpreters/pkg/scanner"
)

var start, current, line int = 0, 0, 0

func main() {
	// scanner.ScanTokens()
	if len(os.Args) < 2 {
		runPrompt()
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		fmt.Println(err)
	} else {
		fmt.Println("Usage: go run main.go [file]")
	}
}

func runPrompt() error {
	fmt.Println("Run prompt!")

	reader := bufio.NewReader(os.Stdin)

	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				run(s)
				return nil
			}
			return err
		}
		run(s)
	}

	return nil
}

func runFile(filePath string) error {
	fmt.Println("Run file!")

	_, err := os.Open(filePath)
	if err != nil {
		return err
	}

	return nil
}

func run(source string) {
	fmt.Println("running run!")
}
