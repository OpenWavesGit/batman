package foldy

import (
	"fmt"
	"os"
)

func Foldy(nam string) {
	// Name of the folder you want to create
	folderName := nam

	// Use "." to represent the current directory
	path := "."

	// Concatenate folder name and path to get the full directory path
	folderPath := path + "/" + folderName

	// Use os.MkdirAll to create the folder and any necessary parent directories
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	fmt.Println("Folder created successfully at", folderPath)
}
