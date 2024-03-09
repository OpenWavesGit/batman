package aesfile

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// EncryptFile encrypts the contents of inputFile and writes the encrypted data to outputFile using AES encryption with the provided key.
func EncryptFile(inputFile, outputFile string, key []byte) error {
	// Open input file
	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inputFileHandle.Close()

	// Create output file
	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFileHandle.Close()

	// Generate a new AES cipher using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new GCM (Galois Counter Mode) cipher using AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generate a random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Write nonce to the beginning of the output file
	if _, err = outputFileHandle.Write(nonce); err != nil {
		return err
	}

	// Create a stream cipher using GCM mode with the nonce
	stream := cipher.NewCTR(block, nonce)

	// Create a writer that writes encrypted data to the output file
	writer := &cipher.StreamWriter{S: stream, W: outputFileHandle}

	// Copy data from input file to output file, encrypting as we go
	if _, err = io.Copy(writer, inputFileHandle); err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts the contents of inputFile and writes the decrypted data to outputFile using AES decryption with the provided key.
func DecryptFile(inputFile, outputFile string, key []byte) error {
	// Open input file
	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inputFileHandle.Close()

	// Read the nonce from the beginning of the input file
	nonce := make([]byte, 12) // Nonce size for GCM mode is 12 bytes
	if _, err := io.ReadFull(inputFileHandle, nonce); err != nil {
		return err
	}

	// Create a new AES cipher using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a stream cipher using GCM mode with the nonce
	stream := cipher.NewCTR(block, nonce)

	// Create a reader that reads encrypted data from the input file
	reader := &cipher.StreamReader{S: stream, R: inputFileHandle}

	// Create output file
	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFileHandle.Close()

	// Copy data from input file to output file, decrypting as we go
	if _, err := io.Copy(outputFileHandle, reader); err != nil {
		return err
	}

	return nil
}
