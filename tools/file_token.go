package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/gtank/cryptopasta"
)

func FileToken(path string, password string, saltHex string) string {
	realPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		panic("failed to resolve symlink: " + err.Error())
	}

	data, err := os.ReadFile(realPath)
	if err != nil {
		panic("failed to read file: " + err.Error())
	}

	sum := sha256.Sum256(data)
	key := DeriveKey(password, saltHex)

	ciphertext, err := cryptopasta.Encrypt(sum[:], key)
	if err != nil {
		panic("failed to encrypt: " + err.Error())
	}

	encoded := hex.EncodeToString(ciphertext)
	return encoded
}
