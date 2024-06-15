package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Cut utility ===

Receives STDIN, splits by delimiter (TAB) into columns, displays the requested columns.

Support flags:
-f - "fields" - select fields (columns)
-d - "delimiter" - use a different delimiter
-s - "separated" - only lines with a separator

The program must pass all tests. The code must pass go vet and golint checks.
*/

var (
	fieldFlag     = flag.String("f", "0", "which columns to print")
	delimiterFlag = flag.String("d", "\t", "delimiter for splitting lines")
	separatedFlag = flag.Bool("s", false, "show only lines with delimiter")
)

// parseColumns parses the column indices from a string
func parseColumns(query string) ([]int, error) {
	numberStrings := strings.Split(strings.ReplaceAll(query, ", ", ","), ",")
	var indexes []int
	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, number)
	}
	return indexes, nil
}

// Cut processes the input data according to the specified flags
func Cut(data string, f string, d string, s bool) (string, error) {
	cols, err := parseColumns(f)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	for _, line := range lines {
		if strings.Contains(line, d) {
			lineSlice := strings.Split(line, d)
			for i, index := range cols {
				if index >= len(lineSlice) {
					continue
				}
				if i > 0 {
					result.WriteString(" ")
				}
				result.WriteString(lineSlice[index])
			}
			result.WriteString("\n")
		} else if !s {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}

func main() {
	// Init
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide a file to read from")
		os.Exit(1)
	}
	src := args[0]
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("No such file")
		os.Exit(1)
	}

	// Call Cut
	result, err := Cut(string(data), *fieldFlag, *delimiterFlag, *separatedFlag)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print(result)
}
