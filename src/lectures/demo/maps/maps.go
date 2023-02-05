package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["Eggs"] = 11
	shoppingList["Bread"] += 1
	shoppingList["Milk"] = 8

	shoppingList["Eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "Milk")
	fmt.Println(shoppingList)

	fmt.Println("Need ", shoppingList["Eggs"], "Eggs")

	_, found := shoppingList["Cereal"]

	if !found {
		fmt.Println("Cereal not found at the shopping list")
	} else {
		fmt.Println("Yeah, Cereal is on")
	}

	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total Items on the list", totalItems)
}
