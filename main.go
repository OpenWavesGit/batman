// main.go
package main

import (
	"fmt"
	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/batrunner"
	"github.com/OpenWavesGit/batman/terminal"
	"os"
	"bufio"
	"github.com/OpenWavesGit/batman/uniquekey"
	"github.com/OpenWavesGit/batman/readers"

	
)

func keygen() {
    keyLength := 10
    filename := "generated_key.txt"
    err := uniquekey.GenerateKey(keyLength, filename)
    if err != nil {
        fmt.Println("Error generating key:", err)
        return
    }
    fmt.Println("Key generated and stored in:", filename)
}



func main() {

//
reader := bufio.NewReader(os.Stdin)

fmt.Println("EVER AUTOMATION!")




fmt.Printf("Hello,  Welcome to the Ever Automation Console App!\n")
//
fmt.Printf("Enter the Location of the Batch file!\n")	
    inputFile := readers.Readfile()

	encryptedFile := "encrypted.txt"
	decryptedFile := "hello.bat"

	// Encrypt the input file
	err := aesfile.EncryptFile(inputFile, encryptedFile)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Decrypt the encrypted file
	err = aesfile.DecodeFile(encryptedFile, decryptedFile)
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
	
	fmt.Println("Press Enter to exit")
	
	reader.ReadString('\n')
	terminal.DrawTerm()
}
