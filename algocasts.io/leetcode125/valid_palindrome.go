package leetcode125

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	a := []rune{}
	for _, r := range s {
		if isDigest(r) || isLetter(r) {
			a = append(a, r)
		}
	}
	i, j := 0, len(a)-1
	for {
		if i >= j {
			return true
		}
		if a[i] != a[j] {
			return false
		} else {
			i++
			j--
		}
	}

	return false
}

func isDigest(s rune) bool {
	if unicode.IsDigit(s) {
		return true
	}
	return false
}

func isLetter(s rune) bool {
	if unicode.IsLetter(s) {
		return true
	}
	return false
}
