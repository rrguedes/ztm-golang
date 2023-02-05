package main

import "fmt"

func main() {
	var myName = "Rafael"
	fmt.Println("My name is", myName)

	var name string = "Cibelle"
	fmt.Println("O nome da minha esposa é ", name)

	userName := "Admin"
	fmt.Println("User Name: ", userName)

	var sum int
	fmt.Println("A soma é: ", sum)

	part1, other := 1, 5
	fmt.Println("A parte 1 e a outra parte", part1, other)

	part2, other := 1, 0
	fmt.Println("A parte 1 e a outra parte", part2, other)

	sum = part1 + part2
	println("A soma total é ", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println(lessonName, lessonType)

	word1, word2, _ := "Hello", "World", "!"

	println("Frase: ", word1, word2)

}
