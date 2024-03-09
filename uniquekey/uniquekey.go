package uniquekey

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// GenerateKey generates a unique key of the specified length and appends it to a file in the specified folder.
func GenerateKey(length int, folderName, fileName string) error {
	key := make([]byte, length)
	for i := range key {
		key[i] = charset[seededRand.Intn(len(charset))]
	}

	// Create folder if it doesn't exist
	if err := os.MkdirAll(folderName, 0755); err != nil {
		return err
	}

	// Open file in append mode or create it if it doesn't exist
	filePath := filepath.Join(folderName, fileName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append key to the file
	if _, err := file.Write(key); err != nil {
		return err
	}

	// Append newline for better readability
	if _, err := file.WriteString("\n"); err != nil {
		return err
	}

	return nil
}

func keygen(kname string) {
	keyLength := 10
	filename := kname + ".txt"
	err := GenerateKey(keyLength, kname, filename)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

}
