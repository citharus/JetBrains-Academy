package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Scan(&num1, &num2, &num3)

	mean := (num1 + num2 + num3) / 3.0

	fmt.Println(mean)
	fmt.Println("Congratulations, you are accepted!")
}
