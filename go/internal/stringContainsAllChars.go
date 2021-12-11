package internal

import "strings"

// returns true if all characters in string s2 are in string s1
func StringContainsAllChars(s1, s2 string) bool {
	for _, c := range s2 {
		if !strings.ContainsRune(s1, c) {
			return false
		}
	}
	return true
}
