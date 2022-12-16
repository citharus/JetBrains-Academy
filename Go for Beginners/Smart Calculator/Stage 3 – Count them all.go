package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

		var sum int
		for _, s := range input {
			number, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}

		fmt.Println(sum)
	}
} 
