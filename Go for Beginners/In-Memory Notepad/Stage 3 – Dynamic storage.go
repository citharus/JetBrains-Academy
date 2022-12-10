package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter the maximum number of notes: ")

	var max int
	fmt.Scanf("%d", &max)

	var records []string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command and data: ")
		scanner.Scan()

		input := strings.SplitN(scanner.Text(), " ", 2)
		command := input[0]

		if command == "exit" {
			fmt.Println("[Info] Bye!")
			break
		} else if command == "create" {
			if len(input) < 2 {
				fmt.Println("[Error] Missing note argument")
				continue
			}
			data := input[1]
			if len(records) < max {
				records = append(records, data)
			} else {
				fmt.Println("[Error] Notepad is full")
				continue
			}
			fmt.Println("[OK] The note was successfully created")
		} else if command == "clear" {
			records = nil
			fmt.Println("[OK] All notes were successfully deleted")
		} else if command == "list" {
			if len(records) <= 0 {
				fmt.Println("[Info] Notepad is empty")
			}

			for index, element := range records {
				fmt.Printf("[Info] %d: %s\n", index+1, element)
			}
		} else {
			fmt.Println("[Error] Unknown command")
		}
	}
}