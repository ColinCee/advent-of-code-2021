package internal

import (
	"bufio"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// continue if line is empty
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
