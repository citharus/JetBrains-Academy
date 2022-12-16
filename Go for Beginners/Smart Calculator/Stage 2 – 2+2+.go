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
		} else if input[0] == "/exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		if len(input) == 1 {
			fmt.Println(input[0])
			continue
		}

		x, err := strconv.Atoi(input[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}


		fmt.Println(x + y)
	}
}