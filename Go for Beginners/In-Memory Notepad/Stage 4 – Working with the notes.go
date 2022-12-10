package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var records []string
var max int

func createCommand(input []string) {
	if len(input) < 2 {
		err := errors.New("[Error] Missing note argument")
		fmt.Println(err.Error())
		return
	}
	data := input[1]
	if len(records) < max {
		records = append(records, data)
	} else {
		err := errors.New("[Error] Notepad is full")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("[OK] The note was successfully created")
}

func updateCommand(input []string) {
	if len(input) < 2 {
		err := errors.New("[Error] Missing position argument")
		fmt.Println(err.Error())
		return
	}
	data := strings.SplitN(input[1], " ", 2)
	if len(data) < 2 {
		err := errors.New("[Error] Missing note argument")
		fmt.Println(err.Error())
		return
	}
	position, err := strconv.Atoi(data[0])
	if err != nil {
		err = fmt.Errorf("[Error] Invalid position: %s", data[0])
		fmt.Println(err.Error())
		return
	}

	if position > max {
		err := fmt.Errorf("[Error] Position %d is out of the boundary [1, %d]", position, max)
		fmt.Println(err.Error())
		return
	} else if position > len(records) {
		err := errors.New("[Error] There is nothing to update")
		fmt.Println(err.Error())
		return
	}

	records[position-1] = data[1]
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
}

func deleteCommand(input []string) {
	if len(input) < 2 {
		err := errors.New("[Error] Missing position argument")
		fmt.Println(err.Error())
		return
	}
	position, err := strconv.Atoi(input[1])
	if err != nil {
		err = fmt.Errorf("[Error] Invalid position: %s", input[1])
		fmt.Println(err.Error())
		return
	}

	if position > max {
		err := fmt.Errorf("[Error] Position %d is out of the boundary [1, %d]", position, max)
		fmt.Println(err.Error())
		return
	} else if position > len(records) {
		err := errors.New("[Error] There is nothing to delete")
		fmt.Println(err.Error())
		return
	}

	records = append(records[:position-1], records[position:]...)
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
}

func clearCommand() {
	records = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

func listCommand() {
	if len(records) <= 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for index, element := range records {
		fmt.Printf("[Info] %d: %s\n", index+1, element)
	}
}

func main() {
	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scanf("%d", &max)

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
			createCommand(input)
		} else if command == "update" {
			updateCommand(input)
		} else if command == "delete" {
			deleteCommand(input)
		} else if command == "clear" {
			clearCommand()
		} else if command == "list" {
			listCommand()
		} else {
			fmt.Println("[Error] Unknown command")
		}
	}
}