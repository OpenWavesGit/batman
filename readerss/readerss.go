package readerss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readd() string {
	// Prompt the user to input the file path

	// Read the input from the user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.TrimSpace(scanner.Text())
	return a
}

// HashTable represents a hash table
type HashTable map[string]string

// LoadHashes loads hashes from a file into a hash table
func LoadHashes(filename string) (HashTable, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hashTable := make(HashTable)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hash := strings.TrimSpace(line)
		hashTable[hash] = "" // Assuming there is no associated value in the file
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hashTable, nil
}

// RemoveHash removes a hash from the file
func RemoveHash(filename, hashToRemove string) error {
	// Load all hashes from the file
	hashes, err := LoadHashes(filename)
	if err != nil {
		return err
	}

	// Remove the hash
	delete(hashes, hashToRemove)

	// Write back the remaining hashes to the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for hash := range hashes {
		fmt.Fprintln(file, hash)
	}

	return nil
}