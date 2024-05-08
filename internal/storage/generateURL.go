package storage

import (
	"crypto/sha1"
	"encoding/base64"
)

func generateShortURL(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	hashBytes := hash.Sum(nil)

	shortCode := base64.URLEncoding.EncodeToString(hashBytes)

	shortCode = shortCode[:8]

	return shortCode
}

func GetURLShortener() func(string)string {
	return generateShortURL
}
