package tools

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"time"

	"github.com/gtank/cryptopasta"
)

func VerifySelf(password string, saltHex string, token string, magicIndex *uint8) uint8 {
	magicNumber := uint8(time.Now().UnixNano()%255) + 1
	execPath, err := os.Executable()
	if err != nil {
		panic("failed to get executable path: " + err.Error())
	}

	realPath, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		panic("failed to resolve symlink: " + err.Error())
	}

	data, err := os.ReadFile(realPath)
	if err != nil {
		panic("failed to read self file: " + err.Error())
	}

	sum := sha256.Sum256(data)
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
		panic("token does not match the file hash, file may have been tampered")
	}
	*magicIndex = *magicIndex ^ magicNumber
	return magicNumber
}
