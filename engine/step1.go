package engine

import (
	"strings"
)

func Step1(message string) string {

	words := strings.Split(message, " ")

	var result []string

	for i, word := range words {

		i = i + 4

		if i%4 == 0 {
			result = append(result, pigLatin(word))
		}

		if i%4 == 1 {
			result = append(result, reverse(word))
		}

		if i%4 == 2 {
			lengthOfWord := len(word)
			secondHalf := word[lengthOfWord/2:]
			firstHalf := word[:lengthOfWord/2]

			result = append(result, secondHalf+firstHalf)
		}

		if i%4 == 3 {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

func reverse(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func pigLatin(word string) string {

	result := []rune{}

	runes := []rune(word)

	firstLetter := runes[0:1]
	otherLetters := runes[1:]

	result = append(otherLetters, firstLetter[0])
	result = append(result, 'a')

	return string(result)
}
