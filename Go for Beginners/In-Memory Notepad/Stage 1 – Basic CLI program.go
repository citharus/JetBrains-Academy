 package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Enter a command and data:")
		wordScanner := bufio.NewScanner(os.Stdin)
		wordScanner.Split(bufio.ScanWords)

		wordScanner.Scan()

		if wordScanner.Text() == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		}

		fmt.Println(wordScanner.Text())

	}
}
