package main

import "fmt"

func main() {
	var name string

	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")

	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
}
