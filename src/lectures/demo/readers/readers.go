package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		input, err := r.ReadString('\n')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading Stdin ", err)
		}
	}
	fmt.Println("Sum ", sum)
}
