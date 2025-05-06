package tools

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gtank/cryptopasta"
)

func VerifyString(input string, password string, saltHex string, token string) {
	sum := sha256.Sum256([]byte(input))
	key := DeriveKey(password, saltHex)

	ciphertext, err := hex.DecodeString(token)
	if err != nil {
		panic("token format error (not hex): " + err.Error())
	}

	plaintext, err := cryptopasta.Decrypt(ciphertext, key)
	if err != nil {
		panic("failed to decrypt token, file may have been tampered: " + err.Error())
	}

	if !bytes.Equal(sum[:], plaintext) {
		panic("token does not match the string hash, file may have been tampered")
	}
}
