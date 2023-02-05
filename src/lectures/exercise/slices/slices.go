//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func main() {

	assemblyLine := []Part{"Part1", "Part2", "Part3"}
	fmt.Println("--- Step 1 ---")
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Part4", "Part5")
	fmt.Println("--- Step 2 ---")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Println("--- Step 3 ---")
	showLine(assemblyLine)
}

func showLine(assemblyLine []Part) {
	for i := 0; i < len(assemblyLine); i++ {
		current := assemblyLine[i]
		fmt.Println(current)
	}
}
