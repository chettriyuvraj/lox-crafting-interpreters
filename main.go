package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/chettriyuvraj/lox-crafting-interpreters/pkg/scanner"
)

/* var hadError bool = false - might not need this because of Go's multiple return types*/

func main() {
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

	sc := scanner.Scanner{}

	tokens, err := sc.ScanTokens()
	if err != nil {
		return err
	}

	fmt.Println(tokens)

	return nil
}

func run(source string) error {
	fmt.Println("running run!")
	sc := scanner.Scanner{}

	tokens, err := sc.ScanTokens()
	if err != nil {
		return err
	}

	fmt.Println(tokens)

	return nil
}

func handleError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[Line %d] Error %s: %s", line, where, message)
}
