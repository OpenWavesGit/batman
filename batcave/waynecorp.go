// main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/batrunner"
	"github.com/OpenWavesGit/batman/filehasher"
	"github.com/OpenWavesGit/batman/readerss"
)

func main1() {

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
	fmt.Println("Processed successfully.")

	// Path to the batch file
	batchFilePath := "hello.bat"

	// Execute the batch file using the package
	output, err := batrunner.RunBatchFile(batchFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output of the batch file
	fmt.Println("OUTPUT = ",output)

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
		fmt.Println("Error file:", err)
		return
	}

	// If there was no error, print a success message
	//fmt.Println("success")
}

func main() {

	fmt.Print("Enter password: ")
	inputs := readerss.Readd()
	hinput := filehasher.FindHash(inputs)
		filename := "bats.hash"

	hashTable, err := readerss.LoadHashes(filename)
	if err != nil {
		fmt.Println("Error loading hashes:", err)
		return
	}



	if _, found := hashTable[hinput]; found {
		fmt.Println("Password Correct")

		// Remove the found hash from the file
		err := readerss.RemoveHash(filename, hinput)
		if err != nil {
			fmt.Println("_", err)
			return
		}
		main1()
		
	} else {
		fmt.Println("Password not Correct")
	}

}
