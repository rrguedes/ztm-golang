//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var totalGeral int = 0

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	totalByFile := func(fileName string) {
		file, err := os.Open(fileName)
		fmt.Printf("Arquivo %v \n", fileName)
		if err != nil {
			fmt.Printf("Não foi possível abrir o arquivo")
		}
		fileLines := bufio.NewScanner(file)
		total := 0
		defer file.Close()
		for fileLines.Scan() {
			// fmt.Println(fileLines.Text())
			number, err := strconv.Atoi(strings.TrimSpace(fileLines.Text()))
			if err == nil {
				total += number
			}
		}
		fmt.Printf("Total do arquivo: %v \n", total)
		totalGeral += total
		time.Sleep(100 * time.Millisecond)
	}
	// Chamar arquivos
	for i := 0; i < len(files); i++ {
		fileName := files[i]
		go totalByFile(fileName)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("--- Total Geral: ", totalGeral)
}
