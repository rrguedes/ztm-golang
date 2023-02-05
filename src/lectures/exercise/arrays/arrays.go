//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func main() {
	var shoppingList [4]Product

	celular := Product{name: "Celular", price: 2908.58}
	bateria := Product{name: "Bateria", price: 908.58}
	capa := Product{name: "Capa", price: 128.99}

	shoppingList[0] = celular
	shoppingList[1] = bateria
	shoppingList[2] = capa

	println("--- BEFORE adding 4th Element ---")
	println("Last item on the list is", getLastElement(shoppingList[:]).name)
	println("Total items on the list", len(shoppingList))
	println("Total Cost is", fmt.Sprintf("%.2f", totalCost(shoppingList[:])))

	shoppingList[3] = Product{name: "Notebook", price: 26000.99}
	println("--- AFTER adding 4th Element ---")
	println("Last item on the list is", getLastElement(shoppingList[:]).name)
	println("Total items on the list", len(shoppingList))
	println("Total Cost is", fmt.Sprintf("%.2f", totalCost(shoppingList[:])))
}

func totalCost(shoppingList []Product) float32 {
	var totalCost float32
	for i := 0; i < len(shoppingList); i++ {
		current := shoppingList[i]
		totalCost += totalCost + current.price
	}
	return totalCost
}

func getLastElement(shoppingList []Product) Product {
	lastElement := shoppingList[len(shoppingList)-1]
	return lastElement
}
