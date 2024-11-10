package core

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Keygen(size int) []byte {
	key := make([]byte, size)
	io.ReadFull(rand.Reader, key)
	return key
}

func Passgen(size int) string {
	return base64.URLEncoding.EncodeToString(Keygen(size))
}
