package internal

import "strconv"

func StringArrayToIntArray(s []string) []int {
	var output []int
	for _, v := range s {
		integer, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		output = append(output, integer)
	}
	return output
}
