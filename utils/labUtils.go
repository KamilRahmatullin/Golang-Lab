package utils

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ExchangeNumbers(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
}

func CheckStringLength(s string, min, max int) bool {
	if utf8.RuneCountInString(s) < min || utf8.RuneCountInString(s) > max {
		return false
	} else {
		return true
	}
}

func IsPalindrome(word string) bool {
	word = strings.ToLower(word)

	tempWord := ""

	for _, char := range word {
		tempWord = string(char) + tempWord
	}

	if tempWord == word {
		return true
	} else {
		return false
	}
}

func IsAnagram(word1, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)
	set1 := make(map[string]int)
	set2 := make(map[string]int)

	for _, char := range word1 {
		set1[string(char)] = set1[string(char)] + 1
	}

	for _, char := range word2 {
		set2[string(char)] = set2[string(char)] + 1
	}

	if reflect.DeepEqual(set1, set2) {
		return true
	} else {
		return false
	}
}

func Reverse(str *string) {
	var newString string

	for _, char := range *str {
		newString = string(char) + newString
	}

	*str = newString
}

func WordsCount(str string) int {
	total := 1
	for _, char := range str {
		if unicode.IsSpace(char) {
			total++
		}
	}

	return total
}

func FindLongestWord(str string) string {
	var longestWord string = ""

	words := strings.Split(str, " ")

	for _, word := range words {
		word = strings.TrimSuffix(word, ",")
		word = strings.TrimSuffix(word, ".")
		word = strings.TrimSuffix(word, "!")
		word = strings.TrimSuffix(word, "?")

		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return longestWord
}

func ClearString(str string) string {
	newS := ""

	for _, ch := range str {
		if unicode.IsLetter(ch) {
			newS = newS + string(ch)
		}
	}

	return newS
}
