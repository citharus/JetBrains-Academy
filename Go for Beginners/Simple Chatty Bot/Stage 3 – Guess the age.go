package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	fmt.Println("Please, remind me your name.")

	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var remainder3 int
	var remainder5 int
	var remainder7 int

	fmt.Print("> ")
	fmt.Scan(&remainder3)
	fmt.Print("> ")
	fmt.Scan(&remainder5)
	fmt.Print("> ")
	fmt.Scan(&remainder7)

	var age int = (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}
