package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func getNumber(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func getResult(input []string) int {
	var result int
	var currentOperation rune
	for _, s := range input {
		if isOperator(s) {
			currentOperation = getOperator(s)
			continue
		}

		switch currentOperation {
		case '-': result -= getNumber(s)
		//case '+': sum += getNumber(s)
		default: result += getNumber(s)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())

		if len(input) == 0 {
			continue
		} else if input[0] == "/help" {
			fmt.Println("The program calculates the sum of numbers")
			continue
		} else if input[0] == "/exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		result := getResult(input)
		fmt.Println(result)
	}
}