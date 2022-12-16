package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isOperator(s string) bool {
	if len(s) == 0 {
		return false
	}

	var count int
	for _, r := range s {
		if r == '-' || r == '+' {
			count++
		}
	}

	return len(s) == count
}

func getOperator(s string) rune {
	if len(s) == 0 || len(s) % 2 == 0 || strings.Contains(s, "+") {
		return '+'
	}
	return '-'
}

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		s = s[1:]
	}

	var count int
	for _, r := range s {
		if unicode.IsNumber(r) {
			count++
		}
	}
	return len(s) == count
}

func getNumber(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func calculateResult(input []string) (int, error) {
	var result int
	var currentOperation rune
	for _, s := range input {
		if isOperator(s) {
			currentOperation = getOperator(s)
			continue
		}

		if !isNumber(s) {
			return 0, errors.New("NaN")
		}

		switch currentOperation {
		case '-': result -= getNumber(s)
		//case '+': result += getNumber(s)
		default: result += getNumber(s)
		}
	}

	if currentOperation == 0 && len(input) > 1 {
		return 0, errors.New("missing operator")
	}

	return result, nil
}

func handleCommands(s string) error {
	s = strings.TrimPrefix(s, "/")

	switch s {
	case "help":
		fmt.Println("The program calculates the sum of numbers")
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		return errors.New("unknown command")
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())

		if len(input) == 0 {
			continue
		}

		if strings.HasPrefix(input[0], "/") {
			err := handleCommands(input[0])
			if err != nil {
				fmt.Println("Unknown command")
			}
			continue
		}

		result, err := calculateResult(input)
		if err != nil {
			fmt.Println("Invalid expression")
			continue
		}
		fmt.Println(result)
	}
}