package terminal

import (
    "os"
    "golang.org/x/term"
	"fmt"
)

// GetTerminalSize returns the width and height of the terminal window.
func GetTerminalSize() (width, height int, err error) {
    // Get the file descriptor for standard output
    fd := int(os.Stdout.Fd())

    // Get the terminal size
    width, height, err = term.GetSize(fd)
    return
}

func DrawTerm(){
width, _, err := GetTerminalSize()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	numSpaces := width - 11
	for i := 0; i < numSpaces; i++ {
		fmt.Print("-")
	}
	fmt.Print("OPENWAVES--")
}