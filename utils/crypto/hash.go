package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
