package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randInt := generateRandInt(1, 100)

	fmt.Println("Generating Inifite Random Numbers")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	randIntCount := generateRandIntn(2, 1, 10)
	fmt.Println("Generating 2 Random Numbers")
	for i := range randIntCount {
		fmt.Println(i)
	}

	fmt.Println("Generating 3 Random Numbers")
	randIntCount2 := generateRandIntn(3, 1, 10)
	// Padrao util para quando buscar dados de um canal fechado retorna valores padrao
	for {
		n, open := <-randIntCount2
		if !open {
			break
		}
		fmt.Println(n)
	}
}
