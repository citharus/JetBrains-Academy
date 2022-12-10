package main

import "fmt"

var money, water, milk, beans, cups int

func buyCommand() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var coffee int
	fmt.Scanf("%d", &coffee)

	if coffee == 1 {
		water -= 250
		beans -= 16
		cups -= 1
		money += 4
	} else if coffee == 2 {
		water -= 350
		milk -= 75
		beans -= 20
		cups -= 1
		money += 7
	} else if coffee == 3 {
		water -= 200
		milk -= 100
		beans -= 12
		cups -= 1
		money += 6
	}
}

func fillCommand() {
	var addWater, addMilk, addBeans, addCups int

	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scanf("%d", &addWater)

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scanf("%d", &addMilk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scanf("%d", &addBeans)

	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scanf("%d", &addCups)

	water += addWater
	milk += addMilk
	beans += addBeans
	cups += addCups
}

func takeCommand() {
	fmt.Printf("I gave you $%d", money)
	money = 0
}

func displayCommand() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n%d ml of milk\n%d g of coffee beans\n%d disposable cups\n$%d of money", water, milk, beans, cups, money)
}

func main() {
	money = 550
	water = 400
	milk = 540
	beans = 120
	cups = 9

	displayCommand()
	fmt.Print("\n\n")

	fmt.Println("Write action (buy, fill, take):")
	var action string
	fmt.Scan(&action)

	if action == "buy" {
		buyCommand()
	} else if action == "fill" {
		fillCommand()
	} else if action == "take" {
		takeCommand()
	}

	fmt.Print("\n\n")
	displayCommand()
}
