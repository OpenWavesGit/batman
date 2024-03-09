// main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/batrunner"
)

func main() {
	
	encryptedFile := "encrypted.enc"
	
	decryptedFile := "hello.bat"
	//
	reader := bufio.NewReader(os.Stdin)

	// Decrypt the encrypted file
	err := aesfile.DecodeFile(encryptedFile, decryptedFile)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return
	}
	fmt.Println("File decrypted successfully.")

	// Path to the batch file
	batchFilePath := "hello.bat"

	// Execute the batch file using the package
	output, err := batrunner.RunBatchFile(batchFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output of the batch file
	fmt.Println(output)

	Deel()
	
	reader.ReadString('\n')
	fmt.Println("Press Enter to exit")

}

func Deel() {
    // Specify the file path
    filePath := "hello.bat"

    // Attempt to remove the file
    err := os.Remove(filePath)
    if err != nil {
        // If there was an error, print it
        fmt.Println("Error deleting file:", err)
        return
    }

    // If there was no error, print a success message
    fmt.Println("File deleted successfully.")
}


// func searchTextInFile(filename string, searchText string) (bool, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.Contains(line, searchText) {
// 			return true, nil
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return false, err
// 	}

// 	return false, nil
// }

// func main() {
// 	filename := "sample.txt"    // Change this to your file name
// 	searchText := "search text" // Change this to the text you want to search for

// 	found, err := searchTextInFile(filename, searchText)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	if found {
// 		fmt.Printf("The text '%s' was found in the file.\n", searchText)
// 	} else {
// 		fmt.Printf("The text '%s' was not found in the file.\n", searchText)
// 	}
// }
