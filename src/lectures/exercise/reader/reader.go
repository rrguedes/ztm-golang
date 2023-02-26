//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandsRan = 0
var nonBlankLines = 0

func main() {
	fmt.Println("Bem vindo champs!")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		trimmedCommand := strings.TrimSpace(input)
		AckCommand(trimmedCommand)
	}
}

func PlusOne() {
	commandsRan += 1
}

func ShowStats() {
	fmt.Println("--- Total commands ran: ---", commandsRan)
	fmt.Println("--- Total non-blank lines: ---", nonBlankLines)
}

func AckCommand(cmd string) {
	responses := map[string]string{
		"hello": "Hello Stranger!",
		"bye":   "Bye bye!",
	}
	switch cmd {
	case "hello":
		fmt.Println(responses["hello"])
		PlusOne()
	case "bye":
		fmt.Println(responses["bye"])
		PlusOne()
	case "q", "Q":
		ShowStats()
	default:
		nonBlankLines += 1
	}
}
