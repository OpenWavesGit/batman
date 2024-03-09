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

func main() {
	var intt int
	encryptedFile := "encrypted.enc"
	encryptedKey := "key.enc"
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
