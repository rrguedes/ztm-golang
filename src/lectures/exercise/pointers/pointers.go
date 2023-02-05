//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	ACTIVE   = true
	INACTIVE = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func main() {
	itemList := []Item{
		{"Item1", ACTIVE},
		{"Item2", ACTIVE},
		{"Item3", ACTIVE},
		{"Item4", ACTIVE},
	}
	fmt.Println(itemList)
	setState(&itemList[2].tag, INACTIVE)
	fmt.Println(itemList)
	checkout(itemList)
	fmt.Println(itemList)
}

func setState(tag *SecurityTag, targetState SecurityTag) {
	*tag = SecurityTag(targetState)
}

func checkout(items []Item) {
	fmt.Println("--- Checkout ---")
	for i := 0; i < len(items); i++ {
		setState(&items[i].tag, INACTIVE)
	}
}
