package internal

import "strings"

func GetDifferenceBetweenTwoStrings(a string, b string) string {
	var result string
	for _, c := range a {
		if !strings.ContainsRune(b, c) {
			result += string(c)
		}
	}
	return result

}
