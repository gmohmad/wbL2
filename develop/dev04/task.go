package main

import (
	"sort"
	"strings"
	"fmt"
)

/*
=== Search for anagrams in the dictionary ===

Write a function to search for all sets of anagrams in a dictionary.
For example:
'пятак', 'пятка' and 'тяпка' - belong to the same set,
'листок', 'слиток' and 'столик' - to another.

Input data for the function: a link to an array - each element of which is a word in Russian in utf8 encoding.
Output: Link to a map of sets of anagrams.
Key - the first word from a set found in the dictionary
The value is a reference to an array, each element of which is a word from the set. The array must be sorted in ascending order.
Sets of one element should not be included in the result.
All words must be converted to lowercase.
As a result, each word should appear only once.

The program must pass all tests. The code must pass go vet and golint checks.
*/

// Helper function to sort characters in a string
func sortString(s string) string {
	sortedChars := strings.Split(s, "")
	sort.Strings(sortedChars)
	return strings.Join(sortedChars, "")
}

// FindAnagramSets function
func FindAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Normalize word (convert to lowercase)
		normalized := strings.ToLower(word)

		// Sort letters in the word
		sortedChars := sortString(normalized)

		// Add to the corresponding anagram set
		anagramSets[sortedChars] = append(anagramSets[sortedChars], normalized)
	}

	// Remove sets with only one element
	for key, set := range anagramSets {
		if len(set) == 1 {
			delete(anagramSets, key)
		} else {
			// Sort the set in ascending order
			sort.Strings(set)
			anagramSets[key] = set
		}
	}

	return anagramSets
}

func main() {
	// Example usage
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "мама", "амам", "рама"}

	anagramSets := FindAnagramSets(words)

	// Print out the result
	for key, set := range anagramSets {
		fmt.Printf("%s: %v\n", key, set)
	}
}
