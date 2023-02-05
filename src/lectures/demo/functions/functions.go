package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello World!")
}

func main() {
	greet()
	fmt.Println(double(2))
	fmt.Println(add(2, 6))
}
