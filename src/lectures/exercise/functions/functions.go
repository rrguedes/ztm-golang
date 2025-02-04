//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func greet(name string) {
	fmt.Println("Hello,", name, "how are you?")
}

func city() string {
	return "Belo Horizonte"
}

func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func getRandomNumber() int {
	return rand.Intn(100)
}

func getTwoRandomNumbers() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func main() {
	greet("Rafael")
	fmt.Println("City Name:", city())
	fmt.Println("Soma dos números 1, 2 e 3, igual a", add(1, 2, 3))
	println(getRandomNumber())
	number1, number2 := getTwoRandomNumbers()
	fmt.Println("Two Random Numbers", number1, number2)
}
