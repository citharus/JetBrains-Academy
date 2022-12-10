package main

import (
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

	var word string
	_, err = fmt.Scan(&word)
	if err != nil {
		log.Fatal(err)
	}
	word = strings.ToLower(word)

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	tabooWords := strings.Split(string(file), "\n")

	for _, w := range tabooWords {
		if w == word {
			fmt.Println("True")
			return
		}
	}
	fmt.Println("False")
}
