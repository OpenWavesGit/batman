// uniquekey/generator.go
package uniquekey

import (
    "io/ioutil"
    "math/rand"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
    rand.NewSource(time.Now().UnixNano()))

// GenerateKey generates a unique key of the specified length and saves it to a file.
func GenerateKey(length int, filename string) error {
    key := make([]byte, length)
    for i := range key {
        key[i] = charset[seededRand.Intn(len(charset))]
    }

    err := ioutil.WriteFile(filename, key, 0644)
    if err != nil {
        return err
    }

    return nil
}
