//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// Estrutura de dados que carrega o Mutex e o contador
type LetterCounter struct {
	totalLetters int
	sync.Mutex
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func countLetters(wg *sync.WaitGroup, letterCounter *LetterCounter, currentWord string) {
	wait()
	letterCounter.Lock()
	defer letterCounter.Unlock()
	defer wg.Done()
	letterCounter.totalLetters += len(currentWord)
}

func main() {
	//--Summary:
	//  Create a program that can read text from standard input and count the
	//  number of letters present in the input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-- Word Counter: Type some sentences so that I can count the words --")
	data, _ := reader.ReadString('\n')
	var wg sync.WaitGroup
	listOfWords := strings.Split(data, " ")
	letterCounter := LetterCounter{}
	for i := 0; i < len(listOfWords); i++ {
		currentWord := listOfWords[i]
		wg.Add(1)
		go countLetters(&wg, &letterCounter, currentWord)
	}
	fmt.Println("Waiting for goroutines")
	wg.Wait()
	letterCounter.Lock()
	totalLetters := letterCounter.totalLetters
	letterCounter.Unlock()
	fmt.Println("Total letters:", totalLetters)
	//--Requirements:
	//* Count the total number of letters in any chosen input
	//* The input must be supplied from standard input
	//* Input analysis must occur per-word, and each word must be analyzed
	//  within a goroutine

	//* When the program finishes, display the total number of letters counted
	//
	//--Notes:
	//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
	//  entering data
	//* Use `cat FILE | go run ./exercise/sync` to analyze a file
	//* Use any synchronization techniques to implement the program:
	//  - Channels / mutexes / wait groups
}
