package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var path string
	_, err := fmt.Scan(&path)
	if err != nil {
		log.Fatal(err)
	}

	tabooWords, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(tabooWords))
}
