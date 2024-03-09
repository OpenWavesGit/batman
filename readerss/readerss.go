package readerss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readd() string {
	// Prompt the user to input the file path

	// Read the input from the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.TrimSpace(scanner.Text())
	return a
}

func FindAndDeleteString(filename, searchString string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			continue // Skip this line
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	// Write back lines to the file
	if err := file.Truncate(0); err != nil {
		return false, err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return false, err
	}
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return len(lines) < len(scanner.Text()), writer.Flush()
}
