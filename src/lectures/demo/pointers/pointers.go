package main

import "fmt"

type Counter struct {
	hits int
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}

func increment(counter *Counter) {
	// Não precisa * quando for struct
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}
