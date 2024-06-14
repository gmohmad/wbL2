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
	var result strings.Builder
	runes := []rune(s)
	escaped := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		// Handle escape sequence
		if char == '\\' && !escaped {
			if i+1 < len(runes) && (runes[i+1] == '\\' || unicode.IsDigit(runes[i+1])) {
				escaped = true
				continue
			} else {
				return "", errors.New("invalid string")
			}
		}

		if unicode.IsDigit(char) && !escaped {
			// Invalid string if the string starts with a digit
			if i == 0 {
				return "", errors.New("invalid string")
			}
			// Get the previous character to repeat
			prevChar := runes[i-1]
			count, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			// Append prevChar count-1 times because it is already appended once
			result.WriteString(strings.Repeat(string(prevChar), count-1))
		} else {
			// Append the character to the result
			result.WriteRune(char)
		}
		escaped = false
	}

	return result.String(), nil
}
