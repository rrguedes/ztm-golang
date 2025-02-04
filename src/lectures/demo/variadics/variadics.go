package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	all := append(a, b...)
	answer := sum(all...)
	fmt.Println(answer)
}
