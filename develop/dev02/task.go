package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Unpacking task ===

Create a Go function that performs a primitive unpacking of a string containing repeated characters/runes, for example:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (incorrect line)
  - "" => ""
Additional task: support for escape sequences
  - qwe\4\5 => qwe45 (*)
  - qwe\45 ​​=> qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

If an incorrect string was passed, the function should return an error. Write unit tests.

The function must pass all tests. The code must pass go vet and golint checks.
*/

func unpackString(s string) (string, error) {
	s += "!" // ! to the string so the last character of the string would be properly handled
	var result strings.Builder
	var currChar rune

	currNum := []rune{}
	escaped := false

	for i, char := range s {
		// Handle escape character
		if char == '\\' && !escaped {
			escaped = true
			continue
		}

		// Handle numbers
		if unicode.IsDigit(char) && !escaped {
			// If the first character of the string is a number, it's invalid
			if i == 0 {
				return "", errors.New("invalid string")
			}
			currNum = append(currNum, char)
			continue
		}
		// Handle other characters

		var num int = 1
		var err error

		if len(currNum) > 0 {
			num, err = strconv.Atoi(string(currNum))
			if err != nil {
				return "", err
			}
		}
		result.WriteString(strings.Repeat(string(currChar), num))

		currChar = char
		currNum = []rune{}
		escaped = false
	}

	// Return result[1:] because on the first iteration we add '0' rune to result
	return result.String()[1:], nil
}
