package tools

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(password string, saltHex string) *[32]byte {
	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		panic("failed to decode salt: " + err.Error())
	}

	key := pbkdf2.Key([]byte(password), salt, 100_000, 32, sha256.New)
	var key32 [32]byte
	copy(key32[:], key)
	return &key32
}
