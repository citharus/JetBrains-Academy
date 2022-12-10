package main

import "fmt"

func main() {
	fmt.Println("Write how many cups of coffee you will need:")
	var amount int
	fmt.Scanf("%d", &amount)
	fmt.Println("For 25 cups of coffee you will need:")
	fmt.Printf("%d ml of water\n", 200*amount)
	fmt.Printf("%d ml of milk\n", 50*amount)
	fmt.Printf("%d g of coffee beans", 15*amount)
}