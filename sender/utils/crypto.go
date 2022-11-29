package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenHash(salt []byte) string {
	h := sha256.New()
	h.Write(salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
