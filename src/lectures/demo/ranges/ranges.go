package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, element := range slice {
		fmt.Println(i, element, ":")
		// Cada letra do elemento
		for _, chr := range element {
			fmt.Printf("	%q\n", chr)
		}
	}
}
