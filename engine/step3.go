package engine

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var crazyWeirdness [26]rune = [26]rune{
	'-', '|', '\\', '/', 'Д', 'Б', 'Є', 'Г', 'Я', '=', 'Ξ', 'Җ', 'Җ',
	'-', '|', '\\', '/', 'Д', 'Б', 'Є', 'Г', 'Я', '=', 'Ξ', 'Җ', 'Җ',
}

func Step3(message string) string {

	words := strings.Split(message, " ")

	var result []string

	for _, word := range words {
		crazyWord := crazifyWord(word)
		result = append(result, crazyWord)
	}

	return strings.Join(result, ",")
}

func crazifyWord(word string) string {

	var result []rune

	for _, letter := range word {
		crazyRune := crazifyRune(letter)
		result = append(result, crazyRune)
	}

	return string(result)
}

func crazifyRune(r rune) rune {

	upperCaseRune := unicode.ToUpper(r)

	coolGuy, _ := utf8.DecodeRuneInString(string(upperCaseRune))

	return crazyWeirdness[coolGuy-65]
}
