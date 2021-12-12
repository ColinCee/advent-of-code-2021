package internal

import "fmt"

func PrintRune(r rune) {
	fmt.Printf("%c", r)
}
func PrintRuneArray(r []rune) {
	for _, v := range r {
		PrintRune(v)
	}
	fmt.Println("")
}
