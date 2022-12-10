package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var path string
	_, err := fmt.Scan(&path)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	tabooWords := strings.Split(string(file), "\n")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		word := scanner.Text()

		if word == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		var isTaboo bool
		for _, w := range tabooWords {
			if w == strings.ToLower(word) {
				isTaboo = true
				break
			}
		}

		if isTaboo == true {
			fmt.Println(strings.Repeat("*", len(word)))
		} else {
			fmt.Println(word)
		}
	}
}
