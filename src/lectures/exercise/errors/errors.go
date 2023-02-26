//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// * Example time string: 14:07:33
func ParseTime(timeString string) (Time, error) {
	timeComponents := strings.Split(timeString, ":")
	if len(timeComponents) != 3 {
		return Time{}, &TimeParseError{"Quantidade de componentes inválida", timeString}
	} else {
		hour, err := strconv.Atoi(timeComponents[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("A hora não pode ser convertida para inteiro: %v", err), timeString}
		}
		minutes, err := strconv.Atoi(timeComponents[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Os minutos não podem ser convertida para inteiro: %v", err), timeString}
		}
		seconds, err := strconv.Atoi(timeComponents[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Os segundos não podem ser convertida para inteiro: %v", err), timeString}
		}
		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("A hora está fora dos ranges requeridos > 23 ou < 0: %v", err), timeString}
		}
		if minutes > 59 || minutes < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("Os minutos estão fora dos ranges requeridos > 59 ou < 0: %v", err), timeString}
		}
		if seconds > 59 || seconds < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("Os segundos estão fora dos ranges requeridos > 59 ou < 0: %v", err), timeString}
		}
		return Time{hour, minutes, seconds}, nil
	}
}
