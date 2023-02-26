//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func NumberOfDigits(data []string) int {
	total := 0
	for i := 0; i < len(data); i++ {
		lineData := data[i]
		runes := []rune(lineData)
		for j := 0; j < len(runes); j++ {
			if unicode.IsDigit(runes[j]) {
				total += 1
			}
		}
	}
	return total
}

func NumberOfLetters(data []string) int {
	total := 0
	for i := 0; i < len(data); i++ {
		lineData := data[i]
		runes := []rune(lineData)
		for j := 0; j < len(runes); j++ {
			if unicode.IsLetter(runes[j]) {
				total += 1
			}
		}
	}
	return total
}

func NumberOfSpaces(data []string) int {
	total := 0
	for i := 0; i < len(data); i++ {
		lineData := data[i]
		runes := []rune(lineData)
		for j := 0; j < len(runes); j++ {
			if unicode.IsSpace(runes[j]) {
				total += 1
			}
		}
	}
	return total
}

func NumberOfPonctuationMarks(data []string) int {
	total := 0
	for i := 0; i < len(data); i++ {
		lineData := data[i]
		runes := []rune(lineData)
		for j := 0; j < len(runes); j++ {
			if unicode.IsPunct(runes[j]) {
				total += 1
			}
		}
	}
	return total
}

func report(data []string, operation func(data []string) int, opName string) {
	fmt.Printf("There are %v %v at the provided string sequence \n", operation(data), opName)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	report(lines, NumberOfDigits, "digits")
	report(lines, NumberOfLetters, "letters")
	report(lines, NumberOfSpaces, "spaces")
	report(lines, NumberOfPonctuationMarks, "ponctuation marks")
}
