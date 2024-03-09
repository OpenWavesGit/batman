package readerss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readd() string {
	// Prompt the user to input the file path
	fmt.Printf("Enter the Location of the Batch file!\n")
	// Read the input from the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.TrimSpace(scanner.Text())
	return a
}
