package Authentication

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func GetKey(int) string {
	key, err := GenerateKey(16)
	if err != nil {
		log.Fatal(err)
	}
	return key
}
