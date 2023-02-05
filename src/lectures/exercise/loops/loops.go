//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	//* Print integers 1 to 50, except:
	//  - Print "Fizz" if the integer is divisible by 3
	//  - Print "Buzz" if the integer is divisible by 5
	//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5

	for i := 0; i < 50; i++ {

		switch cond := i; {
		case cond%3 == 0:
			fmt.Println("Fizz", cond)
		case cond%5 == 0:
			fmt.Println("Buzz", cond)
		case cond%5 == 0 && cond%3 == 0:
			fmt.Println("Fizz Buzz", cond)
		}
	}

}
