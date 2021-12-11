package day8

import (
	"advent-of-code-2021/internal"
	"fmt"
	"strconv"
	"strings"
)

func fourDiffOneSet(segments []string) string {
	four := segments[4]
	one := segments[1]

	// return the difference between the two strings
	return internal.GetDifferenceBetweenTwoStrings(four, one)
}

func sevenDiffOneSet(segments []string) string {
	seven := segments[7]
	one := segments[1]

	// return the difference between the two strings
	return internal.GetDifferenceBetweenTwoStrings(seven, one)
}

func getInitUniqueSegments(segments []string) []string {
	// intialise segmentValues with 10 empty strings
	segmentValues := []string{}
	for i := 0; i < 10; i++ {
		segmentValues = append(segmentValues, "")
	}

	// find the string with length 2 in input
	for _, segment := range segments {
		if len(segment) == 2 {
			segmentValues[1] = segment
		}

		if len(segment) == 3 {
			segmentValues[7] = segment
		}

		if len(segment) == 4 {
			segmentValues[4] = segment
		}

		if len(segment) == 7 {
			segmentValues[8] = segment
		}
	}

	return segmentValues
}

func findFiveLengthSegments(input []string, segments []string) {
	// create and add to set the strings with length 5 in input
	fiveLengthStrings := make(map[string]bool)
	for _, segment := range input {
		if len(segment) == 5 {
			fiveLengthStrings[segment] = true
		}
	}

	fourDiffOne := fourDiffOneSet(segments)
	sevenDiffOne := sevenDiffOneSet(segments)

	// find 3
	for segment, _ := range fiveLengthStrings {
		if internal.StringContainsAllChars(segment, segments[1]) {
			segments[3] = segment
		}
	}

	// remove segment[3] from fiveLengthStrings
	delete(fiveLengthStrings, segments[3])

	// find 5
	for segment, _ := range fiveLengthStrings {
		if internal.StringContainsAllChars(segment, fourDiffOne) && internal.StringContainsAllChars(segment, sevenDiffOne) {
			segments[5] = segment
		}
	}

	// remove segment[5] from fiveLengthStrings
	delete(fiveLengthStrings, segments[5])

	// find 2
	for segment, _ := range fiveLengthStrings {
		segments[2] = segment
	}
}

func findSixLengthSegments(input []string, segments []string) {
	// create and add to set the strings with length 5 in input
	sixLengthStrings := make(map[string]bool)
	for _, segment := range input {
		if len(segment) == 6 {
			sixLengthStrings[segment] = true
		}
	}
	fmt.Println("sixLengthStrings: ", sixLengthStrings)

	sevenDiffOne := sevenDiffOneSet(segments)

	// find nine
	for segment, _ := range sixLengthStrings {
		if internal.StringContainsAllChars(segment, segments[4]) && internal.StringContainsAllChars(segment, sevenDiffOne) {
			segments[9] = segment
		}
	}

	// remove segment[9] from sixLengthStrings
	delete(sixLengthStrings, segments[9])

	// find the last two numbers (0 and 6)
	for segment, _ := range sixLengthStrings {
		if internal.StringContainsAllChars(segment, segments[1]) {
			segments[0] = segment
		}
	}

	delete(sixLengthStrings, segments[0])
	// get first element of sixLengthStrings
	for segment, _ := range sixLengthStrings {
		segments[6] = segment
	}
}

func decodeSegments(line string) int {
	inputOutput := strings.Split(line, "|")
	input := strings.Split(inputOutput[0], " ")
	output := strings.Split(inputOutput[1], " ")

	segments := getInitUniqueSegments(input)
	fmt.Println("segments: ", segments)
	findFiveLengthSegments(input, segments)
	fmt.Println("segments: ", segments)
	findSixLengthSegments(input, segments)
	fmt.Println("segments: ", segments)

	segmentMap := make(map[string]int)
	for index, segment := range segments {
		segmentMap[segment] = index
	}
	fmt.Println("segmentMap: ", segmentMap)

	totalString := ""
	for _, segment := range output {
		for segmentString, segmentValue := range segmentMap {
			if len(segment) == len(segmentString) && internal.StringContainsAllChars(segment, segmentString) {
				number := segmentValue
				fmt.Println("segment: ", segment)
				fmt.Println("number: ", number)
				totalString += strconv.Itoa(number)
			}
		}
	}

	fmt.Println("totalString: ", totalString)
	// convert totalString to int
	total, err := strconv.Atoi(totalString)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return total
}

func SolvePartTwo() {
	lines, err := internal.ReadLines("internal/day8/real_input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range lines {
		total += decodeSegments(line)
	}
	fmt.Println("total in output: ", total)
}
