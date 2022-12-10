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
	tabooWords := strings.Fields(string(file))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		sentence := strings.TrimSuffix(input, ".")

		for _, w := range strings.Fields(sentence) {
			for _, t := range tabooWords {
				if strings.EqualFold(w, t) {
					sentence = strings.ReplaceAll(sentence, w, strings.Repeat("*", len(w)))
				}
			}
		}

		fmt.Print(sentence)
		if strings.ContainsAny(input, ".") {
			fmt.Println(".")
		}
	}
}
