// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/batrunner"
)

func main1() {
	encryptedFile := "encrypted.enc"

	decryptedFile := "hello.bat"
	//
	reader := bufio.NewReader(os.Stdin)

	// Decrypt the encrypted file
	err1 := aesfile.DecodeFile(encryptedFile, decryptedFile)
	if err1 != nil {
		fmt.Println("Error decrypting file:", err1)
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
	reader.ReadString('\n')
	fmt.Println("Press Enter to exit")

}

func searchTextInFile(filename string, searchText string) (bool, error) {
	
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchText) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func main() {
	var pass string
	filename := "key.enc" 
	fmt.Println("enter the password")
	fmt.Scanln(pass)

	
	found, err := searchTextInFile(filename, pass)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if found {
		fmt.Printf("The text '%s' was found in the file.\n", pass)
	} else {
		fmt.Printf("The text '%s' was not found in the file.\n", pass)
	}

	main1()

}
