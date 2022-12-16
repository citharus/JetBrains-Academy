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

type Variables map[string]int

func (variables *Variables) getVariable(name string) (int, bool) {
	for k, v := range *variables {
		if k == name {
			return v, true
		}
	}
	return 0, false
}

func (variables *Variables) setVariable(assignment []string) error {
	name := strings.TrimSpace(assignment[0])
	if !variables.isValidIdentifier(name) {
		return errors.New("Invalid identifier")
	}

	rawValue := strings.TrimSpace(assignment[1])
	var value int
	if !variables.isValidIdentifier(rawValue) && !isNumber(rawValue){
		return errors.New("Invalid assignment")
	}

	if val, ok := variables.getVariable(rawValue); ok && !isNumber(rawValue){
		value = val
	} else if isNumber(rawValue) {
		value = getNumber(rawValue)
	} else {
		return errors.New("Unknown variable")
	}

	(*variables)[name] = value
	return nil
}

func (variables *Variables) isValidIdentifier(identifier string) bool {
	if len(identifier) == 0 {return false}

	var count int
	for _, r := range identifier {
		if unicode.IsLetter(r) {
			count++
		}
	}
	return count == len(identifier)
}

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
	if len(s) == 0 || len(s)%2 == 0 || strings.Contains(s, "+") {
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

func calculateResult(variables *Variables, input []string) (int, error) {
	var result int
	var currentOperation rune
	for _, s := range input {
		if isOperator(s) {
			currentOperation = getOperator(s)
			continue
		}

		if !isNumber(s) {
			if !variables.isValidIdentifier(s) {
				return 0, errors.New("Invalid identifier")
			} else if _, ok := variables.getVariable(s); !ok {
				return 0, errors.New("Unknown variable")
			}
		}

		var number int
		if val, ok := variables.getVariable(s); ok {
			number = val
		} else {
			number = getNumber(s)
		}

		switch currentOperation {
		case '-':
			result -= number
		//case '+': result += getNumber(s)
		default:
			result += number
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
	var variables = make(Variables)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		if strings.HasPrefix(input, "/") {
			err := handleCommands(input)
			if err != nil {
				fmt.Println("Unknown command")
			}
			continue
		}

		if strings.Contains(input, "=") {
			err := variables.setVariable(strings.SplitN(input, "=", 2))
			if err != nil {
				fmt.Println(err)
				continue
			}
			continue
		}

		result, err := calculateResult(&variables, strings.Fields(input))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}