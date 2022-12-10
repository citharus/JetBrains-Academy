package main

import (
	"fmt"
)

func main() {
	var water int
	var milk int
	var beans int
	var amount int

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scanf("%d", &water)

	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scanf("%d", &milk)

	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scanf("%d", &beans)

	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scanf("%d", &amount)

	ingredients := []int{water / 200, milk / 50, beans / 15}
	min := ingredients[0]
	for _, v := range ingredients {
		if min > v {
			min = v
		}
	}

	if min == amount {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if min > amount {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", min-amount)
	} else {
		fmt.Printf("No, I can make only %d cups of coffee\n", min)
	}
}
