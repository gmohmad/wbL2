package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Sort utility ===

Sort rows (man sort)
Basics

Support keys

-k — specifying the column to sort
-n - sort by numeric value
-r - sort in reverse order
-u - do not print duplicate lines

Additional

Support keys

-M - sort by month name
-b - ignore trailing spaces
-c — check if the data is sorted
-h — sort by numerical value, taking into account suffixes

The program must pass all tests. The code must pass go vet and golint checks.
*/

var reverse = flag.Bool("r", false, "Sort in reversed order")
var numerically = flag.Bool("n", false, "Sort numerically")
var unique = flag.Bool("u", false, "Dont print duplicates")
var column = flag.Int("k", 0, "Order by specified column")

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := make([]string, 0)
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Sort function
func Sort(data []byte, r, n, u bool, k int) (string, error) {
	rows := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	var result string

	// unique flag was set
	if u {
		rows = removeDuplicateStr(rows)
	}

	if n {
		numbers := make([]int, 0)
		for _, row := range rows {
			if numRow, err := strconv.Atoi(row); err == nil {
				numbers = append(numbers, numRow)
			} else {
				return "", errors.New("not numerical data")
			}
		}
		sort.Ints(numbers)
		if r {
			sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		}
		for _, row := range numbers {
			result += fmt.Sprintln(row)
		}
		return result, nil
	}

	rowsOfSlices := make([][]string, 0) // Create a 2D-slice, so we can sort rows by specified column
	for _, row := range rows {
		rowSlice := strings.Split(row, " ")
		rowsOfSlices = append(rowsOfSlices, rowSlice)
	}

	if k < 0 || k >= len(rowsOfSlices[0]) {
		return "", fmt.Errorf("incorrect column number: %d", k)
	}

	sort.Slice(rowsOfSlices, func(i, j int) bool { // Sort lexicographically with account of specified column
		for x := k; x < len(rowsOfSlices[i]); x++ {
			if rowsOfSlices[i][k] == rowsOfSlices[j][k] {
				continue
			}
			if r {
				return rowsOfSlices[i][k] > rowsOfSlices[j][k]
			}

			return rowsOfSlices[i][k] < rowsOfSlices[j][k]
		}
		return true
	})

	for _, rowSlice := range rowsOfSlices { // Print out
		for i := 0; i < len(rowSlice); i++ {
			result += rowSlice[i]
			if i != len(rowSlice)-1 { // if last word of row
				result += " "
			} else {
				result += "\n"
			}
		}
	}

	return result, nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	src := args[0]
	r := *reverse
	n := *numerically
	u := *unique
	k := *column
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	fmt.Print(MySort(data, r, n, u, k))
}
