package aesfile

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptFile(foldername, inputFile, outputFile string) error {
	// Read the file content
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Encrypt the file content
	encryptedData, err := Encrypt(data)
	if err != nil {
		return err
	}

	fn := foldername + "copy"

	// Check if the folder already exists
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		// Folder does not exist, create it
		err := os.Mkdir(fn, 0755) // 0755 is the Unix permission mode
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return err
		}

	}

	filepath := filepath.Join(fn, outputFile)
	err = os.WriteFile(filepath, encryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))

	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func DecodeFile(inputFile, outputFile string) error {
	// Read the encrypted file content
	encryptedData, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Decrypt the file content
	decryptedData, err := Decrypt(encryptedData)
	if err != nil {
		return err
	}

	// Write the decrypted content to a new file
	err = os.WriteFile(outputFile, decryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

