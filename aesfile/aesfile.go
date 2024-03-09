package aesfile

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "os"
)

const (
    MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
)

func Encode(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}

func EncryptFile(inputFile, outputFile string) error {
    // Read the file content
    data, err := os.ReadFile(inputFile)
    if err != nil {
        return err
    }

    // Encrypt the file content
    encryptedData, err := Encrypt(data)
    if err != nil {
        return err
    }

    // Write the encrypted content to a new file
    err = os.WriteFile(outputFile, encryptedData, 0644)
    if err != nil {
        return err
    }

    return nil
}

func Encrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher([]byte(MySecret))
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))

    iv := ciphertext[:aes.BlockSize]
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

    return ciphertext, nil
}
