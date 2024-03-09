// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/OpenWavesGit/batman/aesfile"
	"github.com/OpenWavesGit/batman/foldy"
	"github.com/OpenWavesGit/batman/readerss"
	"github.com/OpenWavesGit/batman/terminal"
	"github.com/OpenWavesGit/batman/uniquekey"
)

// func fold(nam string) {
// 	// Name of the folder you want to create
// 	folderName := nam

// 	// Use "." to represent the current directory
// 	path := "."

// 	// Concatenate folder name and path to get the full directory path
// 	folderPath := path + "/" + folderName

// 	// Use os.MkdirAll to create the folder and any necessary parent directories
// 	err := os.MkdirAll(folderPath, os.ModePerm)
// 	if err != nil {
// 		fmt.Println("Error creating folder:", err)
// 		return
// 	}

// 	fmt.Println("Folder created successfully at", folderPath)
// }

func main() {
	var intt int
	//
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("EVER AUTOMATION!")

	fmt.Printf("Hello,  Welcome to the Ever Automation Batman App!\n")
	//
	fmt.Println("\n enter the company name")
	cmpname := readerss.Readd()
	fmt.Printf("Enter the Location of the Batch file!\n")
	inputFile := readerss.Readd()
	encryptedFile := "encrypted.enc"
	encryptedKey := "key.enc"
	fmt.Println("Enter the Total number of keys for the company")
	fmt.Scanln(&intt)
	foldy(cmpname)
	for i := 1; i <= intt; i++ {
		uniquekey.Keygen(cmpname)
	}
	fmt.Println("Key generated and stored in:", cmpname)
	//	decryptedFile := "hello.bat"

	// Encrypt the input file
	err := aesfile.EncryptFile(cmpname, inputFile, encryptedFile)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Encrypt the input file

	inputfile2 := cmpname + ".txt"
	filepath := filepath.Join(cmpname, inputfile2)
	err1 := aesfile.EncryptFile(cmpname, filepath, encryptedKey)
	if err1 != nil {
		fmt.Println("Error encrypting file:", err1)
		return
	}
	fmt.Println("File encrypted successfully.")

	fmt.Println("Press Enter to exit")
	reader.ReadString('\n')
	terminal.DrawTerm()
}
