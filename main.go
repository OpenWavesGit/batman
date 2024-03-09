// main.go
package main

import (
    "fmt"
    "github.com/OpenWavesGit/batman/batrunner" 
	
	"github.com/OpenWavesGit/batman/aesfile"
)

func main() {

//
//aes
inputFile := "input.txt"
	encryptedFile := "output.enc"
	decryptedFile := "output.txt"
	key := []byte("0123456789abcdef0123456789abcdef") // 32 bytes for AES-256

	// Encrypt file
	if err := aesfile.EncryptFile(inputFile, encryptedFile, key); err != nil {
		fmt.Println("Encryption Error:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Decrypt file
	if err := aesfile.DecryptFile(encryptedFile, decryptedFile, key); err != nil {
		fmt.Println("Decryption Error:", err)
		return
	}
	fmt.Println("File decrypted successfully.")
//aes
//
    // Path to the batch file
    batchFilePath := "path/to/your/batchfile.bat"

    // Execute the batch file using the package
    output, err := batrunner.RunBatchFile(batchFilePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the output of the batch file
    fmt.Println(output)
}
