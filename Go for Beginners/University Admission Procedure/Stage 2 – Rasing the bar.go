package main

import "fmt"

const threshold = 60.0

func main() {
	var num1, num2, num3 float64
	fmt.Scan(&num1, &num2, &num3)

	mean := (num1 + num2 + num3) / 3.0
	fmt.Println(mean)

	if mean >= threshold {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}