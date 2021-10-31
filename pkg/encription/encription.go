package encription

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
)

const blockSize int = aes.BlockSize

func generateRandom(len int) ([]byte, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func EncriptStr(s string) string {
	src := []byte(s)

	key, err := generateRandom(blockSize)

	if err != nil {
		return ""
	}
	aesB, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	dst := make([]byte, blockSize)
	hash := md5.Sum(src)

	aesB.Encrypt(dst, hash[:])

	return fmt.Sprintf("%x", dst)
}
