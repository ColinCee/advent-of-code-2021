package internal

import "fmt"

// print array to console
func PrettyPrintIntArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
}

func PrettyPrint2DIntArray(a [][]int) {
	for i := 0; i < len(a); i++ {
		PrettyPrintIntArray(a[i])
	}
	fmt.Printf("\n")
}
