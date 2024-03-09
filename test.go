package main

import (
	"fmt"
	
	"github.com/OpenWavesGit/batman/uniquekey"
	
	
)

func main() {



keyLength := 20
uniqueKey := uniquekey.GenerateTimePrefixedKey(keyLength)
fmt.Println("Generated Key:", uniqueKey)

}