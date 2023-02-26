//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

func (operation Operation) calculate(a, b int) int {
	switch operation {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	case Divide:
		return a / b
	}
	return 0
}

func main() {
	add := Add
	fmt.Println(add.calculate(2, 2)) // = 4
	sub := Subtract
	fmt.Println(sub.calculate(10, 3)) // = 7
	mul := Multiply
	fmt.Println(mul.calculate(3, 3)) // = 9
	div := Divide
	fmt.Println(div.calculate(100, 2)) // = 50
}
