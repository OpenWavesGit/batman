package main

import (
    "fmt"
    "os"
    "golang.org/x/term"
)

func main() {
    fmt.Print("Enter password: ")

    // Disable echoing of input characters
    password, err := term.ReadPassword(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println("Error reading password:", err)
        return
    }

    // Print a newline to move to the next line after password input
    fmt.Println("\nPassword entered:", string(password))
}
