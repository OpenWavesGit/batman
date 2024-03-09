// filehasher/filehasher.go

package filehasher

import (
    "bufio"
    "crypto/sha256"
    "encoding/hex"
    "os"
)

// HashFile hashes each line in a file and saves it to a new file
func HashFile(inputFilePath, outputFilePath string) error {
    // Open the input file
    inputFile, err := os.Open(inputFilePath)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    // Create the output file
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    // Create a new scanner to read the input file line by line
    scanner := bufio.NewScanner(inputFile)

    // Create a new writer to write to the output file
    writer := bufio.NewWriter(outputFile)

    // Iterate through each line in the input file
    for scanner.Scan() {
        line := scanner.Text()

        // Hash the line using SHA-256
        hashedLine := sha256.Sum256([]byte(line))

        // Convert the hashed line to a hexadecimal string
        hashedLineHex := hex.EncodeToString(hashedLine[:])

        // Write the hashed line to the output file
        _, err := writer.WriteString(hashedLineHex + "\n")
        if err != nil {
            return err
        }
    }

    // Flush any buffered data to the output file
    writer.Flush()

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}

func FindHash(input string) string {
    hash := sha256.New()
    hash.Write([]byte(input))
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}