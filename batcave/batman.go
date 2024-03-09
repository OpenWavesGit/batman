// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/filehasher"
	"github.com/OpenWavesGit/batman/foldy"
	"github.com/OpenWavesGit/batman/readerss"
	"github.com/OpenWavesGit/batman/terminal"
	"github.com/OpenWavesGit/batman/uniquekey"
)

func main() {
	var intt int
	encryptedFile := "encrypted.enc"
//	encryptedKey := "key.enc"
	//
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("EVER AUTOMATION!")
	fmt.Printf("Hello,  Welcome to the Ever Automation Batman App!\n")
	fmt.Println("\n enter the company name")
	cmpname := readerss.Readd()
	fmt.Printf("Enter the Location of the Batch file!\n")
	inputFile := readerss.Readd()

	fmt.Println("Enter the Total number of keys for the company")
	fmt.Scanln(&intt)

	foldy.Foldy(cmpname)

	for i := 1; i <= intt; i++ {
		uniquekey.Keygen(cmpname)
	}
	fmt.Println("Key generated and stored in:", cmpname)

	// Encrypt the batch file

	err := aesfile.EncryptFile(cmpname, inputFile, encryptedFile)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Hash the Kay file

	inputFilePath := cmpname + ".txt"
	outputFilePath := "bats.hash"
	ifilepath := filepath.Join(cmpname, inputFilePath)
    ofilepath := filepath.Join(cmpname+"copy", outputFilePath)
	err1 := filehasher.HashFile(ifilepath, ofilepath)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	fmt.Println("Key Hashing completed successfully.")

	fmt.Println("Press Enter to exit")
	reader.ReadString('\n')
	terminal.DrawTerm()
}
