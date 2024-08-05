package iteration

import (
	"strings"
)

func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func removeSpaces(s string) string {
	sl := strings.Fields(s)
	var result string = ""
	for i := 0; i < cap(sl); i++ {
		if sl[i] == "," {
			result = result + sl[i]
		} else {
			result = result + " " + sl[i]
		}
	}
	return result[1:]
}

func isPalindrome(s string) bool {

	sl := strings.Fields(s)

	var sWithoutSpaces string = ""

	for _, el := range sl {
		sWithoutSpaces = sWithoutSpaces + el
	}
	sep := string(sWithoutSpaces[len(sWithoutSpaces)/2])

	before, after, _ := strings.Cut(sWithoutSpaces, sep)
	after = Reverse(after)
	if strings.Compare(before, after) == 0 {
		return true
	} else {
		return false
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
