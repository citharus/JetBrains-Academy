package main

import "fmt"

var money, water, milk, beans, cups int

func checkResources(nWater int, nMilk int, nBeans int, nCups int) bool {
	if nWater > water {
		fmt.Println("Sorry, not enough water!")
		return false
	} else if nMilk > milk {
		fmt.Println("Sorry, not enough milk!")
		return false
	} else if nBeans > beans {
		fmt.Println("Sorry, not enough coffee beans!")
		return false
	} else if nCups > cups {
		fmt.Println("Sorry, not enough disposable cups!")
		return false
	}

	fmt.Println("I have enough resources, making you a coffee!")
	return true
}

func buyCommand() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var coffee int
	fmt.Scanf("%d", &coffee)

	if coffee == 1 && checkResources(250, 0, 16, 1) {
		water -= 250
		beans -= 16
		cups -= 1
		money += 4
	} else if coffee == 2 && checkResources(350, 75, 20, 1) {
		water -= 350
		milk -= 75
		beans -= 20
		cups -= 1
		money += 7
	} else if coffee == 3 && checkResources(200, 100, 12, 1) {
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
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
	fmt.Println()
}

func main() {
	money = 550
	water = 400
	milk = 540
	beans = 120
	cups = 9

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)

		if action == "buy" {
			buyCommand()
		} else if action == "fill" {
			fillCommand()
		} else if action == "take" {
			takeCommand()
		} else if action == "remaining" {
			displayCommand()
		} else if action == "exit" {
			break
		}
	}
}
