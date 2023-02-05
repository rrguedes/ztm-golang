package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 Scored Higher Than Quiz2")
	} else if quiz1 > quiz2 {
		fmt.Println("Quiz2 Scored Higher Than Quiz1")
	} else {
		fmt.Println("Quiz1 and Quiz2 Scored Equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Passed!")
	} else {
		fmt.Println("Not Passed!")
	}
}
