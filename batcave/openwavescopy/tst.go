package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	filename := "bats.hash" // Replace with the name of your file containing hashes
	hashTable, err := LoadHashes(filename)
	if err != nil {
		fmt.Println("Error loading hashes:", err)
		return
	}

	var searchHash string
	fmt.Print("Enter the hash to search for: ")
	fmt.Scanln(&searchHash)

	if _, found := hashTable[searchHash]; found {
		fmt.Println("Hash found!")
	} else {
		fmt.Println("Hash not found")
	}
}
