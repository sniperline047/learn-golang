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
	hours   int
	minutes int
	seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.input, t.msg)
}

func validRange(value int, min int, max int) bool {
	return value >= min && value < max
}

func ParseTime(timestring string) (Time, error) {
	timestringArray := strings.Split(timestring, ":")

	if len(timestringArray) != 3 {
		return Time{}, &TimeParseError{"Invalid timestring", timestring}
	}

	hours, err := strconv.Atoi(timestringArray[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hours %v", err), timestring}
	}

	minutes, err := strconv.Atoi(timestringArray[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minutes %v", err), timestring}
	}

	seconds, err := strconv.Atoi(timestringArray[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing seconds %v", err), timestring}
	}

	if !validRange(hours, 0, 24) {
		return Time{}, &TimeParseError{fmt.Sprintf("hours out of range [0, 24) : %v", hours), timestring}
	}
	if !validRange(minutes, 0, 60) {
		return Time{}, &TimeParseError{fmt.Sprintf("minutes out of range [0, 60) : %v", minutes), timestring}
	}
	if !validRange(seconds, 0, 60) {
		return Time{}, &TimeParseError{fmt.Sprintf("seconds out of range [0, 60) : %v", seconds), timestring}
	}

	parsedTime := Time{hours, minutes, seconds}
	fmt.Println("Parsed time:", parsedTime)

	return parsedTime, nil
}
