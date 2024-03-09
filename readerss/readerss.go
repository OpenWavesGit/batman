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

func FindAndRemoveStringInFile(filename, searchStr string) error {
    // Open the file
    file, err := os.OpenFile(filename, os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Create a temporary file to store the modified content
    tmpFile, err := os.CreateTemp("", "tempfile")
    if err != nil {
        return err
    }
    defer os.Remove(tmpFile.Name())
    defer tmpFile.Close()

    // Create a writer to write to the temporary file
    writer := bufio.NewWriter(tmpFile)

    // Iterate through each line of the file
    for scanner.Scan() {
        line := scanner.Text()

        // Check if the line contains the search string
        if !strings.Contains(line, searchStr) {
            // Write the line to the temporary file if it does not contain the search string
            _, err := writer.WriteString(line + "\n")
            if err != nil {
                return err
            }
        }
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        return err
    }

    // Flush the writer to ensure all data is written to the file
    if err := writer.Flush(); err != nil {
        return err
    }

    // Close the temporary file
    if err := tmpFile.Close(); err != nil {
        return err
    }

    // Remove the original file
    if err := os.Remove(filename); err != nil {
        return err
    }

    // Rename the temporary file to the original file name
    if err := os.Rename(tmpFile.Name(), filename); err != nil {
        return err
    }

    return nil
}