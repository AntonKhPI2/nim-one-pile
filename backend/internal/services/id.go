package services

import (
	"crypto/rand"
	"encoding/hex"
)

func newID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
