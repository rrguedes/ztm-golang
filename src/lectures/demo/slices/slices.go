package main

import "fmt"

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")

	route = route[2:]
	printSlice("Remaining Locations", route)
}

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}
