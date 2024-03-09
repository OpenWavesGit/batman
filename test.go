package main

import (
	"fmt"
	
	"github.com/OpenWavesGit/batman/uniquekey"
	
	
)

func main() {
	keyLength := 10
	filename := "kname"+".txt"
	err := uniquekey.GenerateKey(keyLength, filename)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	fmt.Println("Key generated and stored in:", filename)
}