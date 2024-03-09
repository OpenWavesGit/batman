package readers

import (
    "bufio"
    "fmt"
    "os"
)

func Readfile() (string, error) {
    // Prompt the user to enter the file path
    fmt.Print("Enter the file path: ")
    
    // Create a scanner to read user input from standard input (keyboard)
    scanner := bufio.NewScanner(os.Stdin)
    
    // Read the input from the user
    if scanner.Scan() {
        // Get the file path entered by the user
        inputFile := scanner.Text()
        
        // Open the file
        file, err := os.Open(inputFile)
        if err != nil {
            return "", fmt.Errorf("failed to open file: %w", err)
        }
        defer file.Close()
        
        // Create a scanner to read the file line by line
        fileScanner := bufio.NewScanner(file)
        
        // Read the contents of the file line by line
        var fileContent string
        for fileScanner.Scan() {
            fileContent += fileScanner.Text() + "\n"
        }
        
        // Check for errors during scanning
        if err := fileScanner.Err(); err != nil {
            return "", fmt.Errorf("error reading file: %w", err)
        }
        
        return fileContent, nil

    } else {
        return "", fmt.Errorf("error reading input: %w", scanner.Err())
    } 
}
