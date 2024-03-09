package uniquekey

import (
    "math/rand"
    "time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+={}[]"

var seededRand *rand.Rand = rand.New(
    rand.NewSource(time.Now().UnixNano()))

// GenerateKey generates a unique key of the specified length.
func GenerateKey(length int) string {
    key := make([]byte, length)
    for i := range key {
        key[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(key)
}

