package engine

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Step2(message string) string {

	words := strings.Split(message, " ")

	var result []string

	for _, word := range words {
		swappedWord := swapWord(word)
		result = append(result, swappedWord)
	}

	return strings.Join(result, " ")
}

func swapWord(word string) string {

	var result []rune

	for _, letter := range word {

		swappedLetter := swapLetter(letter)

		result = append(result, swappedLetter)
	}

	return string(result)
}

func swapLetter(r rune) rune {

	upperCaseRune := unicode.ToUpper(r)

	coolGuy, _ := utf8.DecodeRuneInString(string(upperCaseRune))

	if coolGuy >= 78 {
		return coolGuy - 13
	} else {
		return coolGuy + 13
	}
}
